/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/6/15 20:42
 * @Desc: 登录服务器
 */

package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/third/redis"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建redis客户端
	rds := redis.NewClient(
		redis.WithAddrs("127.0.0.1:6379"),
		redis.WithDB(0),
	)

	// 创建grpc服务器
	gs := grpc.NewServer(
		grpc.WithServerListenAddr(":8082"),
	)

	// 创建registry注册发现
	registry := consul.NewRegistry(
		consul.WithAddress("127.0.0.1:8500"),
	)

	// 创建节点服务器
	server := node.NewNode(
		node.WithName("login"),
		node.WithRedis(rds),
		node.WithGRPCServer(gs),
		node.WithRegistry(registry),
	)

	core := NewCore(server.Proxy())
	// 初始化业务
	core.Init()
	// 添加组件
	container.Add(server)
	// 启动服务器
	container.Serve()
}
