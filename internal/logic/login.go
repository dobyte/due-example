package logic

import (
	"fmt"
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due-example/internal/user"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
)

type login struct {
	proxy node.Proxy
	um    *user.Manager
}

func NewLogin(proxy node.Proxy) *login {
	return &login{proxy: proxy, um: user.NewManager()}
}

func (l *login) Init() {
	// 用户注册
	l.proxy.AddRouteHandler(route.Register, false, l.register)
	// 用户登录
	l.proxy.AddRouteHandler(route.Login, false, l.login)
}

// 用户注册
func (l *login) register(r node.Request) {
	req := &pb.RegisterReq{}
	res := &pb.RegisterRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("response register message failed, err: %v", err)
		}
	}()

	fmt.Printf("GID:%s", r.GID())
	fmt.Printf("NID:%s", r.NID())
	fmt.Printf("CID:%d", r.CID())
	fmt.Printf("UID:%d", r.UID())

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid register message, err: %v", err)
		res.Code = pb.RegisterCode_Failed
		return
	}

	if l.um.HasUser(req.Account) {
		res.Code = pb.RegisterCode_AccountExists
		return
	}

	id := l.um.AddUser(&user.User{
		Account:  req.Account,
		Password: req.Password,
		Nickname: req.Nickname,
		Age:      int(req.Age),
	})

	res.Code = pb.RegisterCode_Ok
	res.ID = id
	res.Account = req.Account
}

// 用户登录
func (l *login) login(r node.Request) {
	req := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("response login message failed, err: %v", err)
		}
	}()

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid login message, err: %v", err)
		res.Code = pb.LoginCode_Failed
		return
	}

	u, err := l.um.GetUser(req.Account)
	if err != nil {
		switch err {
		case user.ErrNotFoundUser:
			res.Code = pb.LoginCode_IncorrectAccountOrPassword
		default:
			res.Code = pb.LoginCode_Failed
		}
		return
	}

	if u.Password != req.Password {
		res.Code = pb.LoginCode_IncorrectAccountOrPassword
		return
	}

	if err = r.BindGate(int64(u.ID)); err != nil {
		log.Errorf("bind gate failed: %v", err)
		return
	}

	res.Code = pb.LoginCode_Ok
	res.ID = u.ID
	res.Account = req.Account
}
