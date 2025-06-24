package greet

import (
	greetBiz "Soraka/internal/biz/greet"
	basicReq "Soraka/internal/dal/req/basic"
	"Soraka/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct {
	uc  *greetBiz.GreetUseCase
	win *application.WebviewWindow
}

func NewGreetService() *GreetService {
	return &GreetService{
		uc: greetBiz.NewGreetUseCase(nil),
	}
}

func (s *GreetService) SetMainWin(win *application.WebviewWindow) {
	s.win = win
	s.uc.SetMainWin(win)
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

func (s *GreetService) SetTheme() {
	s.uc.SetTheme()
}
