/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/6/15 20:42
 * @Desc: 登录服务器
 */

package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due-example/distributed-architecture/login-server/business"
	cluster "github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 开启调试模式
	mode.SetMode(mode.DebugMode)
	// 初始化事件总线
	//eventbus.SetEventbus(kafka.NewEventbus())
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter()
	// 创建容器
	container := due.NewContainer()
	// 创建节点组件
	node := cluster.NewNode(
		cluster.WithLocator(locator),
		cluster.WithRegistry(registry),
		cluster.WithTransporter(transporter),
	)
	// 初始化业务
	business.Init(node.Proxy())
	// 添加节点组件
	container.Add(node)
	// 启动容器
	container.Serve()
}
