package router

import (
	"Soraka/internal/service/greet"
	greetWails "Soraka/internal/wails/greet"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type BasicRouter struct {
	greetService *greet.GreetService
}

func NewBasicRouter(
	greetService *greet.GreetService,
) *BasicRouter {
	return &BasicRouter{
		greetService: greetService,
	}
}

func (r *BasicRouter) InitRouter(root *gin.RouterGroup) {
	group := root.Group("/basic")
	r.greetService.RegisterGin(group)
}

func (r *BasicRouter) CollectWailsServices(services *[]application.Service) {
	ws := greetWails.NewGreetSvcWails(r.greetService)
	*services = append(*services, application.NewService(ws))
}

func (r *BasicRouter) SetMainWin(win *application.WebviewWindow) {
	r.greetService.SetMainWin(win)
}
