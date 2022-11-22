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
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := gate.NewGate(
		gate.WithServer(ws.NewServer()),
		gate.WithLocator(redis.NewLocator()),
		gate.WithRegistry(etcd.NewRegistry()),
		gate.WithTransporter(rpcx.NewTransporter()),
	)
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
