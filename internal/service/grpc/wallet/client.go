package wallet

import (
	"github.com/dobyte/due-example/internal/service/grpc/wallet/pb"
	"github.com/dobyte/due/transport"
	"google.golang.org/grpc"
)

const target = "discovery://wallet"

func NewClient(fn transport.NewServiceClientFunc) (pb.WalletClient, error) {
	client, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewWalletClient(client.Client().(*grpc.ClientConn)), nil
}
