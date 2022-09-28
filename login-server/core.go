package main

import (
	"context"

	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due-example/internal/user"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
)

type Core struct {
	um    *user.Manager
	proxy node.Proxy
}

func NewCore(proxy node.Proxy) *Core {
	return &Core{proxy: proxy, um: user.NewManager()}
}

func (c *Core) Init() {
	// 用户注册
	c.proxy.AddRouteHandler(route.Register, false, c.register)
	// 用户登录
	c.proxy.AddRouteHandler(route.Login, false, c.login)
}

// 用户注册
func (c *Core) register(req node.Request) {
	msg := &pb.RegisterReq{}
	res := &pb.RegisterRes{}
	defer func() {
		if err := req.Response(res); err != nil {
			log.Errorf("response register message failed, err: %v", err)
		}
	}()

	if err := req.Parse(msg); err != nil {
		log.Errorf("invalid register message, err: %v", err)
		res.Code = pb.RegisterCode_Failed
		return
	}

	if c.um.HasUser(msg.Account) {
		res.Code = pb.RegisterCode_AccountExists
		return
	}

	id := c.um.AddUser(&user.User{
		Account:  msg.Account,
		Password: msg.Password,
		Nickname: msg.Nickname,
		Age:      int(msg.Age),
	})

	res.Code = pb.RegisterCode_Ok
	res.ID = id
}

// 用户登录
func (c *Core) login(req node.Request) {
	msg := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := req.Response(res); err != nil {
			log.Errorf("response login message failed, err: %v", err)
		}
	}()

	if err := req.Parse(msg); err != nil {
		log.Errorf("invalid login message, err: %v", err)
		res.Code = pb.LoginCode_Failed
		return
	}

	u, err := c.um.GetUser(msg.Account)
	if err != nil {
		switch err {
		case user.ErrNotFoundUser:
			res.Code = pb.LoginCode_IncorrectAccountOrPassword
		default:
			res.Code = pb.LoginCode_Failed
		}
		return
	}

	if u.Password != msg.Password {
		res.Code = pb.LoginCode_IncorrectAccountOrPassword
		return
	}

	gid, cid := req.GID(), req.CID()

	if err = c.proxy.BindGate(context.Background(), gid, cid, int64(u.ID)); err != nil {
		log.Errorf("bind gate failed: %v", err)
		return
	}

	res.Code = pb.LoginCode_Ok
	res.ID = u.ID
}
