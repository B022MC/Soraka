package greet

import (
	greetBiz "Soraka/internal/biz/greet"
	basicReq "Soraka/internal/dal/req/basic"
	"Soraka/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct {
	uc *greetBiz.GreetUseCase
}

func NewGreetService(win *application.WebviewWindow) *GreetService {
	return &GreetService{
		uc: greetBiz.NewGreetUseCase(win),
	}
}

func (s *GreetService) RegisterGin(group gin.IRoutes) {
	group.GET("/hello", s.Hello)
}

func (s *GreetService) Name() string {
	return "GreetService"
}

func (s *GreetService) Hello(ctx *gin.Context) {
	var req basicReq.Hello
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "ok", s.uc.Greet(req.Name))
}

// 可以扩展更多业务方法
func (s *GreetService) SetTheme() {
	s.uc.SetTheme()
}
