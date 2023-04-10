package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due-example/distributed-architecture/mesh-server/business"
	cluster "github.com/dobyte/due/cluster/mesh"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 开启调试模式
	mode.SetMode(mode.DebugMode)
	// 创建容器
	container := due.NewContainer()
	// 创建网格组件
	mesh := cluster.NewMesh(
		cluster.WithLocator(redis.NewLocator()),
		cluster.WithRegistry(consul.NewRegistry()),
		cluster.WithTransporter(rpcx.NewTransporter()),
	)
	// 初始化业务
	business.Init(mesh.Proxy())
	// 添加网格组件
	container.Add(mesh)
	// 启动容器
	container.Serve()
}
