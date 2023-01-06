package route

import (
	"github.com/dobyte/due-example/distributed-architecture/admin-server/app/api"
	"github.com/dobyte/due/cluster/master"
	"github.com/gin-gonic/gin"
)

func Init(proxy master.Proxy, engine *gin.Engine) {
	var (
		notifyAPI = api.NewNotify(proxy)
	)

	// 发送消息
	engine.POST("push", notifyAPI.Push)
	// 组播消息
	engine.POST("multicast", notifyAPI.Multicast)
	// 组播消息
	engine.POST("broadcast", notifyAPI.Broadcast)
}
