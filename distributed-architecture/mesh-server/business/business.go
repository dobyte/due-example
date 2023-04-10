package business

import (
	"github.com/dobyte/due-example/internal/service/rpcx/wallet"
	"github.com/dobyte/due/cluster/mesh"
)

func Init(proxy *mesh.Proxy) {
	// 初始化钱包服务
	wallet.NewServer(proxy).Init()
}
