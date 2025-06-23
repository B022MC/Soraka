package router

import (
	"Soraka/internal/service/greet"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type BasicRouter struct {
	greetService *greet.GreetService
	//lcuApiService *lcu.LcuApiService
}

func NewBasicRouter(
	greetService *greet.GreetService,
	// lcuApiService *lcu.LcuApiService,
) *BasicRouter {
	return &BasicRouter{
		greetService: greetService,
		//lcuApiService: lcuApiService,
	}
}

func (r *BasicRouter) InitRouter(root *gin.RouterGroup) {
	group := root.Group("/basic")
	r.greetService.RegisterGin(group)
	//r.lcuApiService.RegisterGin(group)
}

func (r *BasicRouter) CollectWailsServices(services *[]application.Service) {
	ws := greet.NewWailsService(r.greetService)
	*services = append(*services, application.NewService(ws))
	//*services = append(*services, application.NewService(r.lcuApiService))
}
