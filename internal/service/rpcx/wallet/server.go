package wallet

import (
	"context"
	"github.com/dobyte/due-example/internal/service/rpcx/wallet/pb"
	"github.com/dobyte/due/cluster/mesh"
	"github.com/dobyte/due/log"
)

const (
	service     = "wallet" // 用于客户端定位服务，例如discovery://wallet
	servicePath = "Wallet" // 服务路径要与pb中的服务路径保持一致
)

type Wallet struct {
	proxy *mesh.Proxy
}

func NewServer(proxy *mesh.Proxy) *Wallet {
	return &Wallet{
		proxy: proxy,
	}
}

func (w *Wallet) Init() {
	w.proxy.AddServiceProvider(service, servicePath, w)
}

// IncrCoin 增加金币
func (w *Wallet) IncrCoin(ctx context.Context, req *pb.IncrCoinRequest, reply *pb.IncrCoinReply) error {
	log.Infof("incr %d coin success for uid: %d", req.Coin, req.UID)

	return nil
}

// DecrCoin 减少金币
func (w *Wallet) DecrCoin(ctx context.Context, req *pb.DecrCoinRequest, reply *pb.DecrCoinReply) error {
	log.Infof("decr %d coin success for uid: %d", req.Coin, req.UID)

	return nil
}
