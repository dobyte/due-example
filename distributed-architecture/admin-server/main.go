package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due-example/distributed-architecture/admin-server/app"
	"github.com/dobyte/due/cluster/master"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	m := master.NewMaster(
		master.WithLocator(redis.NewLocator()),
		master.WithRegistry(etcd.NewRegistry()),
		master.WithTransporter(rpcx.NewTransporter()),
	)
	// 创建后台组件
	a := app.NewAdmin(m.Proxy())
	// 添加网关组件
	container.Add(m, a)
	// 启动容器
	container.Serve()
}
