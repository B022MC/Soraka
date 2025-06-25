package router

import (
	"Soraka/internal/service/greet"
	"Soraka/internal/service/lcu"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

type RootRouter struct {
	basicRouter *BasicRouter
	lcuRouter   *LcuRouter
	mainWin     *application.WebviewWindow
}

func NewRootRouter() *RootRouter {
	greetSvc := greet.NewGreetService() // 先不传 window
	basicRouter := NewBasicRouter(greetSvc)

	lcuApiSvc := lcu.NewLcuApiService()
	lcuRouter := NewLcuRouter(lcuApiSvc)

	return &RootRouter{
		basicRouter: basicRouter,
		lcuRouter:   lcuRouter,
	}
}

func (r *RootRouter) SetMainWin(win *application.WebviewWindow) {
	r.mainWin = win
	r.basicRouter.SetMainWin(win)
}

// Gin 注册
func (r *RootRouter) InitRouter(group *gin.RouterGroup) {
	r.basicRouter.InitRouter(group)
	r.lcuRouter.InitRouter(group)
}

// Wails 注册
func (r *RootRouter) CollectWailsServices(services *[]application.Service, log *log.Logger) {
	r.basicRouter.CollectWailsServices(services)
	r.lcuRouter.CollectWailsServices(services, log)
}
