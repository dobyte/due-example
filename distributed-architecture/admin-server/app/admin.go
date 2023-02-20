package app

import (
	"github.com/dobyte/due-example/distributed-architecture/admin-server/app/route"
	"github.com/dobyte/due/cluster/master"
	"github.com/dobyte/due/component"
	"github.com/dobyte/due/config"
	"github.com/dobyte/due/log"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	component.Base
	proxy  *master.Proxy
	engine *gin.Engine
}

func NewAdmin(proxy *master.Proxy) *Admin {
	return &Admin{proxy: proxy, engine: gin.Default()}
}

// Name 组件名称
func (a *Admin) Name() string {
	return "admin"
}

// Init 初始化组件
func (a *Admin) Init() {
	route.Init(a.proxy, a.engine)
}

// Start 启动组件
func (a *Admin) Start() {
	go func() {
		addr := config.Get("config.http.addr", ":8080").String()

		if err := a.engine.Run(addr); err != nil {
			log.Fatal("http server startup failed: %v", err)
		}
	}()
}

// Destroy 销毁组件
func (a *Admin) Destroy() {

}
