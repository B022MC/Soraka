package router

import (
	lcuService "Soraka/internal/service/lcu"
	lcuWails "Soraka/internal/wails/lcu"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

type LcuRouter struct {
	lcuApiService *lcuService.LcuApiService
}

func NewLcuRouter(
	lcuApiService *lcuService.LcuApiService,
) *LcuRouter {
	return &LcuRouter{
		lcuApiService: lcuApiService,
	}
}
func (r *LcuRouter) InitRouter(root *gin.RouterGroup) {
	group := root.Group("/lcu")
	r.lcuApiService.RegisterGin(group)
}
func (r *LcuRouter) CollectWailsServices(services *[]application.Service, log *log.Logger) {
	ws := lcuWails.NewLcuApiWails(r.lcuApiService, log)
	*services = append(*services, application.NewService(ws))
}
