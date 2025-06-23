package router

import (
	"Soraka/internal/service/lcu"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type LcuRouter struct {
	lcuApiService *lcu.LcuApiService
}

func NewLcuRouter(
	lcuApiService *lcu.LcuApiService,
) *LcuRouter {
	return &LcuRouter{
		lcuApiService: lcuApiService,
	}
}
func (r *LcuRouter) InitRouter(root *gin.RouterGroup) {
	group := root.Group("/lcu")
	r.lcuApiService.RegisterGin(group)
}
func (r *LcuRouter) CollectWailsServices(services *[]application.Service) {
	*services = append(*services, application.NewService(r.lcuApiService))
}
