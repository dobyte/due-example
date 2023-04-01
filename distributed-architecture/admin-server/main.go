package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due-example/distributed-architecture/admin-server/app"
	cluster "github.com/dobyte/due/cluster/master"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 开启调试模式
	mode.SetMode(mode.DebugMode)
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter(grpc.WithClientDiscovery(registry))
	// 创建容器
	container := due.NewContainer()
	// 创建管理组件
	master := cluster.NewMaster(
		cluster.WithLocator(locator),
		cluster.WithRegistry(registry),
		cluster.WithTransporter(transporter),
	)
	// 创建后台组件
	a := app.NewAdmin(master.Proxy())
	// 添加组件
	container.Add(master, a)
	// 启动容器
	container.Serve()
}
