package business

import (
	"context"
	"github.com/dobyte/due-example/internal/event"
	"github.com/dobyte/due-example/internal/logic"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/eventbus"
	"github.com/dobyte/due/log"
)

func Init(proxy *node.Proxy) {
	// 初始化路由
	initRoute(proxy)
	// 初始化事件
	initEvent()
}

// 初始化路由
func initRoute(proxy *node.Proxy) {
	// 初始化登录逻辑
	logic.NewLogin(proxy).Init()
}

// 初始化事件
func initEvent() {
	err := eventbus.Subscribe(context.Background(), event.Login, event.LoginHandler)
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}
}
