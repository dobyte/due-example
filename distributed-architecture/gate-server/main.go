/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/7/7 1:53 上午
 * @Desc: 网关服务器
 */

package main

import (
	"github.com/dobyte/due"
	cluster "github.com/dobyte/due/cluster/gate"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/grpc"
	"net/http"
)

func main() {
	// 开启调试模式
	mode.SetMode(mode.DebugMode)
	// 创建容器
	container := due.NewContainer()
	// 创建服务器
	server := ws.NewServer()
	// 监听HTTP连接升级WS协议
	server.OnUpgrade(func(w http.ResponseWriter, r *http.Request) bool {
		return true
	})
	// 创建网关组件
	gate := cluster.NewGate(
		cluster.WithServer(server),
		cluster.WithLocator(redis.NewLocator()),
		cluster.WithRegistry(consul.NewRegistry()),
		cluster.WithTransporter(grpc.NewTransporter()),
	)
	// 添加网关组件
	container.Add(gate)
	// 启动容器
	container.Serve()
}
