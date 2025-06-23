package router

import (
	"Soraka/internal/service/greet"
	"Soraka/internal/service/lcu"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type RootRouter struct {
	basicRouter *BasicRouter
	lcuRouter   *LcuRouter
}

// NewRootRouter 统一构建各模块实例
func NewRootRouter(win *application.WebviewWindow) *RootRouter {
	// 业务服务实例化
	greetSvc := greet.NewGreetService(win)

	// 子路由实例化
	basicRouter := NewBasicRouter(greetSvc)

	lcuApiSvc := lcu.NewLcuApiService()
	lcuRouter := NewLcuRouter(lcuApiSvc)
	return &RootRouter{
		basicRouter: basicRouter,
		lcuRouter:   lcuRouter,
	}
}

// Gin 注册
func (r *RootRouter) InitRouter(group *gin.RouterGroup) {
	r.basicRouter.InitRouter(group)
	r.lcuRouter.InitRouter(group)
}

// Wails 注册
func (r *RootRouter) CollectWailsServices(services *[]application.Service) {
	r.basicRouter.CollectWailsServices(services)
	r.lcuRouter.CollectWailsServices(services)
}
