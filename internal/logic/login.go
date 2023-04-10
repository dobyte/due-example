package logic

import (
	"github.com/dobyte/due-example/internal/event"
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due-example/internal/service/rpcx/wallet"
	walletpb "github.com/dobyte/due-example/internal/service/rpcx/wallet/pb"
	"github.com/dobyte/due-example/internal/user"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/eventbus"
	"github.com/dobyte/due/log"
)

type login struct {
	proxy *node.Proxy
	um    *user.Manager
}

func NewLogin(proxy *node.Proxy) *login {
	return &login{proxy: proxy, um: user.NewManager()}
}

func (l *login) Init() {
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		// 用户注册
		group.AddRouteHandler(route.Register, false, l.register)
		// 用户登录
		group.AddRouteHandler(route.Login, false, l.login)
	})
}

// 用户注册
func (l *login) register(ctx *node.Context) {
	req := &pb.RegisterReq{}
	res := &pb.RegisterRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response register message failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
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
func (l *login) login(ctx *node.Context) {
	req := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response login message failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
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

	if err = ctx.BindGate(int64(u.ID)); err != nil {
		log.Errorf("bind gate failed: %v", err)
		return
	}

	err = eventbus.Publish(ctx.Context(), event.Login, &event.LoginPayload{
		ID:      u.ID,
		Account: u.Account,
	})
	if err != nil {
		log.Errorf("%s event push failed: %v", event.Login, err)
	}

	cli, err := wallet.NewClient(l.proxy.NewServiceClient)
	if err != nil {
		log.Errorf("")
	} else {
		coin := int64(10)
		_, err = cli.IncrCoin(ctx.Context(), &walletpb.IncrCoinRequest{
			UID:  int64(u.ID),
			Coin: coin,
		})
		if err != nil {
			log.Errorf("incr %d coin for %d failed: %v", coin, u.ID, err)
		}
	}

	res.Code = pb.LoginCode_Ok
	res.ID = u.ID
	res.Account = u.Account
}
