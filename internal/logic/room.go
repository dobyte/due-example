package logic

import (
	"github.com/dobyte/due-example/internal/entity"
	"github.com/dobyte/due-example/internal/middleware"
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
)

type room struct {
	rm    *entity.Manager
	proxy *node.Proxy
}

func NewRoom(proxy *node.Proxy) *room {
	return &room{proxy: proxy, rm: entity.NewManager()}
}

func (c *room) Init() {
	// 监听重新连接
	c.proxy.Events().AddEventHandler(cluster.Reconnect, c.reconnect)
	// 监听连接断开
	c.proxy.Events().AddEventHandler(cluster.Disconnect, c.disconnect)
	// 监听路由组
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册中间件
		group.Middleware(middleware.Auth)
		// 创建房间
		group.AddRouteHandler(route.CreateRoom, false, c.createRoom)
		// 加入房间
		group.AddRouteHandler(route.JoinRoom, false, c.joinRoom)
	})
}

// 重新连接
func (c *room) reconnect(evt *node.Event) {
	log.Warnf("connection is reopened, gid: %v, cid: %d, uid: %d", evt.GID, evt.CID, evt.UID)
}

// 连接断开
func (c *room) disconnect(evt *node.Event) {
	log.Warnf("connection is closed, gid: %v, cid: %d, uid: %d", evt.GID, evt.CID, evt.UID)
}

// 创建房间
func (c *room) createRoom(ctx *node.Context) {
	msg := &pb.CreateRoomReq{}
	res := &pb.CreateRoomRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response create room message failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(msg); err != nil {
		log.Errorf("invalid create room message, err: %v", err)
		res.Code = pb.CreateRoomCode_Failed
		return
	}

	if err := ctx.BindNode(); err != nil {
		log.Errorf("bind node failed, err: %v", err)
		res.Code = pb.CreateRoomCode_Failed
		return
	}

	id := c.rm.AddRoom(&entity.Room{
		Name: msg.Name,
	})

	res.Code = pb.CreateRoomCode_Ok
	res.ID = id
}

// 加入房间
func (c *room) joinRoom(ctx *node.Context) {
	// todo
}
