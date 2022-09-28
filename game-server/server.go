package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locator/redis"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/grpc"
)

func init() {
	// 设置模式
	mode.SetMode(mode.DebugMode)

	// 设置日志
	log.SetLogger(log.NewLogger(
		log.WithOutFile("./log/game.log"),
		log.WithCallerSkip(1),
		log.WithOutLevel(log.DebugLevel),
	))
}

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建定位器
	locator := redis.NewLocator(
		redis.WithAddrs("127.0.0.1:6379"),
	)

	// 创建服务发现
	registry := etcd.NewRegistry(
		etcd.WithAddrs("127.0.0.1:2379"),
	)

	// 创建传输器
	transport := grpc.NewServer(
		grpc.WithServerListenAddr(":8083"),
	)

	// 创建网关组件
	component := node.NewNode(
		node.WithName("game"),
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithGRPCServer(transport),
	)

	core := NewCore(component.Proxy())
	// 初始化业务
	core.Init()
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
