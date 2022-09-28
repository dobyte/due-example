/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/7/7 1:53 上午
 * @Desc: 网关服务器
 */

package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/gate"
	"github.com/dobyte/due/locator/redis"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/grpc"
)

func init() {
	// 设置模式
	mode.SetMode(mode.DebugMode)

	// 设置日志
	log.SetLogger(log.NewLogger(
		log.WithOutFile("./log/gate.log"),
		log.WithCallerSkip(1),
		log.WithOutLevel(log.DebugLevel),
	))
}

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建服务器
	server := ws.NewServer(
		ws.WithServerListenAddr(":3553"),
		ws.WithServerMaxConnNum(5000),
	)

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
		grpc.WithServerListenAddr(":8081"),
	)

	// 创建网关组件
	component := gate.NewGate(
		gate.WithName("gate"),
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
		gate.WithGRPCServer(transport),
	)

	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
