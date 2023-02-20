package route

import (
	"github.com/dobyte/due-example/internal/logic"
	"github.com/dobyte/due/cluster/node"
)

func Init(proxy *node.Proxy) {
	// 初始化登录逻辑
	logic.NewLogin(proxy).Init()
}
