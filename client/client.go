package main

import (
	"github.com/dobyte/due"
	_ "github.com/dobyte/due"
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/network/ws"
)

const (
	defaultUserAccount  = "due"
	defaultUserPassword = "123456"
	defaultUserNickname = "fuxiao"
	defaultUserAge      = 31
	defaultRoomName     = "room"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化事件和路由
	initEvent(component.Proxy())
	initRoute(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

func initEvent(proxy client.Proxy) {
	// 打开连接
	proxy.AddEventListener(cluster.Connect, onConnect)
	// 重新连接
	proxy.AddEventListener(cluster.Reconnect, onReconnect)
	// 断开连接
	proxy.AddEventListener(cluster.Disconnect, onDisconnect)
}

func onConnect(proxy client.Proxy) {
	log.Infof("connection is opened")

	err := proxy.Push(0, route.Register, &pb.RegisterReq{
		Account:  defaultUserAccount,
		Password: defaultUserPassword,
		Nickname: defaultUserNickname,
		Age:      defaultUserAge,
	})
	if err != nil {
		log.Errorf("push login message failed: %v", err)
	}

	//err := proxy.Push(0, route.CreateRoom, &pb.CreateRoomReq{
	//	Name: defaultRoomName,
	//})
	//if err != nil {
	//	log.Errorf("push create room message failed: %v", err)
	//}
}

func onReconnect(proxy client.Proxy) {
	log.Infof("connection is reopened")

	err := proxy.Push(0, route.Login, &pb.LoginReq{
		Account:  defaultUserAccount,
		Password: defaultUserPassword,
	})
	if err != nil {
		log.Errorf("push login message failed: %v", err)
	}
}

func onDisconnect(proxy client.Proxy) {
	log.Infof("connection is closed")

	err := proxy.Reconnect()
	if err != nil {
		log.Errorf("reconnect failed: %v", err)
	}
}

func initRoute(proxy client.Proxy) {
	// 用户注册
	proxy.AddRouteHandler(route.Register, registerHandler)
	// 用户登录
	proxy.AddRouteHandler(route.Login, loginHandler)
	// 创建房间
	proxy.AddRouteHandler(route.CreateRoom, createRoomHandler)
	// 通知消息
	proxy.AddRouteHandler(route.NotifyMessage, notifyMessageHandler)
	// 用户未授权
	proxy.AddRouteHandler(route.Unauthorized, unauthorizedHandler)
}

func unauthorizedHandler(r client.Request) {
	log.Errorf("the user request is not authorized")
}

func registerHandler(r client.Request) {
	res := &pb.RegisterRes{}

	err := r.Parse(res)
	if err != nil {
		log.Errorf("invalid register response message, err: %v", err)
		return
	}

	switch res.Code {
	case pb.RegisterCode_Failed:
		log.Error("user register failed")
		return
	case pb.RegisterCode_AccountExists:
		log.Warn("account already exists")
	default:
		log.Infof("user register successful, UserID: %v", res.ID)
	}

	err = r.Proxy().Push(0, route.Login, &pb.LoginReq{
		Account:  defaultUserAccount,
		Password: defaultUserPassword,
	})
	if err != nil {
		log.Errorf("push login message failed: %v", err)
	}
}

func loginHandler(r client.Request) {
	res := &pb.LoginRes{}

	err := r.Parse(res)
	if err != nil {
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

	err = r.Proxy().Push(0, route.CreateRoom, &pb.CreateRoomReq{
		Name: defaultRoomName,
	})
	if err != nil {
		log.Errorf("push create room message failed: %v", err)
	}
}

func createRoomHandler(r client.Request) {
	res := &pb.CreateRoomRes{}

	err := r.Parse(res)
	if err != nil {
		log.Errorf("invalid create room response message, err: %v", err)
		return
	}

	switch res.Code {
	case pb.CreateRoomCode_Failed:
		log.Error("create room failed")
	case pb.CreateRoomCode_NameExists:
		log.Warn("room name already exists")
	default:
		log.Infof("create room successful, RoomID: %v", res.ID)
	}
}

func notifyMessageHandler(r client.Request) {
	msg := &pb.Notify{}

	err := r.Parse(msg)
	if err != nil {
		log.Errorf("invalid notify message, err: %v", err)
		return
	}

	log.Infof("received a message from admin: %s", msg.Message)
}
