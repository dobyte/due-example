package main

import (
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due/encoding"
	"github.com/dobyte/due/encoding/proto"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/network"
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/packet"
)

const (
	defaultUserAccount  = "due"
	defaultUserPassword = "123456"
	defaultUserNickname = "fuxiao"
	defaultUserAge      = 31
	defaultRoomName     = "room"
)

var (
	codec    encoding.Codec
	handlers map[int32]handler
)

type handler func(conn network.Conn, buffer []byte)

func init() {
	codec = encoding.Invoke(proto.Name)
	handlers = map[int32]handler{
		route.Register:   registerHandler,
		route.Login:      loginHandler,
		route.CreateRoom: createRoomHandler,
	}
}

func main() {
	client := ws.NewClient(ws.WithClientDialUrl("ws://127.0.0.1:3553"))

	client.OnConnect(func(conn network.Conn) {
		log.Infof("connection is opened")
	})
	client.OnDisconnect(func(conn network.Conn) {
		log.Infof("connection is closed")
	})
	client.OnReceive(func(conn network.Conn, msg []byte, msgType int) {
		message, err := packet.Unpack(msg)
		if err != nil {
			log.Errorf("unpack message failed: %v", err)
			return
		}

		handler, ok := handlers[message.Route]
		if !ok {
			log.Errorf("the route handler is not registered, route:%v", message.Route)
			return
		}
		handler(conn, message.Buffer)
	})

	conn, err := client.Dial()
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}

	if err = push(conn, route.Register, &pb.RegisterReq{
		Account:  defaultUserAccount,
		Password: defaultUserPassword,
		Nickname: defaultUserNickname,
		Age:      defaultUserAge,
	}); err != nil {
		log.Errorf("push message failed: %v", err)
	}

	select {}
}

func registerHandler(conn network.Conn, buffer []byte) {
	res := &pb.RegisterRes{}
	if err := codec.Unmarshal(buffer, res); err != nil {
		log.Errorf("invalid register response message, err: %v", err)
		return
	}

	switch res.Code {
	case pb.RegisterCode_Failed:
		log.Error("user register failed")
		return
	case pb.RegisterCode_AccountExists:
		log.Error("account already exists")
	default:
		log.Infof("user register successful, UserID: %v", res.ID)
	}

	if err := push(conn, route.Login, &pb.LoginReq{
		Account:  defaultUserAccount,
		Password: defaultUserPassword,
	}); err != nil {
		log.Errorf("push message failed: %v", err)
	}
}

func loginHandler(conn network.Conn, buffer []byte) {
	res := &pb.LoginRes{}
	if err := codec.Unmarshal(buffer, res); err != nil {
		log.Errorf("invalid login response message, err: %v", err)
		return
	}

	switch res.Code {
	case pb.LoginCode_Failed:
		log.Error("user login failed")
		return
	case pb.LoginCode_IncorrectAccountOrPassword:
		log.Error("incorrect account or password")
		return
	default:
		log.Infof("user login successful, UserID: %v", res.ID)
	}

	if err := push(conn, route.CreateRoom, &pb.CreateRoomReq{
		Name: defaultRoomName,
	}); err != nil {
		log.Errorf("push message failed: %v", err)
	}
}

func createRoomHandler(conn network.Conn, buffer []byte) {
	res := &pb.CreateRoomRes{}
	if err := codec.Unmarshal(buffer, res); err != nil {
		log.Errorf("invalid create room response message, err: %v", err)
		return
	}

	switch res.Code {
	case pb.CreateRoomCode_Failed:
		log.Error("create room failed")
		return
	case pb.CreateRoomCode_NameExists:
		log.Error("room name already exists")
		return
	default:
		log.Infof("create room successful, RoomID: %v", res.ID)
	}
}

func push(conn network.Conn, route int32, data interface{}) error {
	buffer, err := codec.Marshal(data)
	if err != nil {
		return err
	}

	msg, err := packet.Pack(&packet.Message{
		Route:  route,
		Buffer: buffer,
	})
	if err != nil {
		return err
	}

	return conn.Push(msg)
}
