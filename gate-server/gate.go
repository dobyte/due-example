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
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/third/redis"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建服务器
	svr := ws.NewServer(
		ws.WithServerListenAddr(":3553"),
		ws.WithServerMaxConnNum(5000),
	)

	// 创建redis客户端
	rds := redis.NewClient(
		redis.WithAddrs("127.0.0.1:6379"),
		redis.WithDB(0),
	)

	// 创建grpc服务器
	gs := grpc.NewServer(
		grpc.WithServerListenAddr(":8081"),
	)

	// 创建registry注册发现
	registry := consul.NewRegistry(
		consul.WithAddress("127.0.0.1:8500"),
	)

	// 创建网关
	server := gate.NewGate(
		gate.WithRedis(rds),
		gate.WithServer(svr),
		gate.WithGRPCServer(gs),
		gate.WithRegistry(registry),
	)

	// 添加组件
	container.Add(server)
	// 启动服务器
	container.Serve()
}
