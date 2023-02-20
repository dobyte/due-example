package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due-example/distributed-architecture/game-server/route"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	n := node.NewNode(
		node.WithLocator(redis.NewLocator()),
		node.WithRegistry(consul.NewRegistry()),
		node.WithTransporter(rpcx.NewTransporter()),
	)
	// 初始化路由
	route.Init(n.Proxy())
	// 添加网关组件
	container.Add(n)
	// 启动容器
	container.Serve()
}
