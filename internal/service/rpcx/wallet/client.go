package wallet

import (
	"github.com/dobyte/due-example/internal/service/rpcx/wallet/pb"
	"github.com/dobyte/due/transport"
	"github.com/smallnest/rpcx/client"
)

const target = "discovery://wallet"

func NewClient(fn transport.NewServiceClientFunc) (*pb.WalletClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewWalletClient(c.Conn().(client.XClient)), nil
}
