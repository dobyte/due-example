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
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/eventbus"
	"github.com/dobyte/due/eventbus/kafka"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 初始化事件总线
	eventbus.SetEventbus(kafka.NewEventbus())
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	n := node.NewNode(
		node.WithLocator(redis.NewLocator()),
		node.WithRegistry(consul.NewRegistry()),
		node.WithTransporter(rpcx.NewTransporter()),
	)
	// 初始化业务
	business.Init(n.Proxy())
	// 添加网关组件
	container.Add(n)
	// 启动容器
	container.Serve()
}
