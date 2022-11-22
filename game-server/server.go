package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := node.NewNode(
		node.WithLocator(redis.NewLocator()),
		node.WithRegistry(etcd.NewRegistry()),
		node.WithTransporter(rpcx.NewTransporter()),
	)
	// 创建业务处理器
	core := NewCore(component.Proxy())
	// 初始化业务
	core.Init()
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
