package main

import (
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/room"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
)

type Core struct {
	rm    *room.Manager
	proxy node.Proxy
}

func NewCore(proxy node.Proxy) *Core {
	return &Core{proxy: proxy, rm: room.NewManager()}
}

func (c *Core) Init() {
	// 监听重新连接
	c.proxy.AddEventListener(cluster.Reconnect, c.reconnect)
	// 监听连接断开
	c.proxy.AddEventListener(cluster.Disconnect, c.disconnect)
	// 创建房间
	c.proxy.AddRoute(route.CreateRoom, false, c.createRoom)
	// 加入房间
	c.proxy.AddRoute(route.JoinRoom, false, c.joinRoom)
}

// 重新连接
func (c *Core) reconnect(gid string, uid int64) {
	log.Warnf("connection is reopened, gid: %v, uid: %d", gid, uid)
}

// 连接断开
func (c *Core) disconnect(gid string, uid int64) {
	log.Warnf("connection is closed, gid: %v, uid: %d", gid, uid)
}

// 创建房间
func (c *Core) createRoom(req node.Request) {
	msg := &pb.CreateRoomReq{}
	res := &pb.CreateRoomRes{}
	defer func() {
		if err := req.Response(res); err != nil {
			log.Errorf("response create room message failed, err: %v", err)
		}
	}()

	if err := req.Parse(msg); err != nil {
		log.Errorf("invalid create room message, err: %v", err)
		res.Code = pb.CreateRoomCode_Failed
		return
	}

	if err := c.proxy.BindNode(req.Context(), req.UID()); err != nil {
		log.Errorf("bind node failed, err: %v", err)
		res.Code = pb.CreateRoomCode_Failed
		return
	}

	id := c.rm.AddRoom(&room.Room{
		Name: msg.Name,
	})

	res.Code = pb.CreateRoomCode_Ok
	res.ID = id
}

// 加入房间
func (c *Core) joinRoom(req node.Request) {

}
