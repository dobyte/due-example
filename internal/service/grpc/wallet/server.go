package wallet

import (
	"context"
	"github.com/dobyte/due-example/internal/service/grpc/wallet/pb"
	"github.com/dobyte/due/cluster/mesh"
	"github.com/dobyte/due/log"
)

const service = "wallet"

type Wallet struct {
	pb.UnimplementedWalletServer
	proxy *mesh.Proxy
}

func NewServer(proxy *mesh.Proxy) *Wallet {
	return &Wallet{proxy: proxy}
}

func (w *Wallet) Init() {
	w.proxy.AddServiceProvider(service, &pb.Wallet_ServiceDesc, w)
}

// IncrCoin 增加金币
func (w *Wallet) IncrCoin(ctx context.Context, req *pb.IncrCoinRequest) (*pb.IncrCoinReply, error) {
	log.Infof("incr %d coin success for uid: %d", req.Coin, req.UID)

	return &pb.IncrCoinReply{}, nil
}

// DecrCoin 减少金币
func (w *Wallet) DecrCoin(ctx context.Context, req *pb.DecrCoinRequest) (*pb.DecrCoinReply, error) {
	log.Infof("decr %d coin success for uid: %d", req.Coin, req.UID)

	return &pb.DecrCoinReply{}, nil
}
