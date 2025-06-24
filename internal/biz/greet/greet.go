package greet

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetUseCase struct {
	win *application.WebviewWindow
}

func NewGreetUseCase(win *application.WebviewWindow) *GreetUseCase {
	return &GreetUseCase{
		win: win,
	}
}

func (uc *GreetUseCase) SetMainWin(win *application.WebviewWindow) {
	uc.win = win
}

// 核心业务逻辑：返回问候语
func (uc *GreetUseCase) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// 核心业务逻辑：设置主题/窗口操作
func (uc *GreetUseCase) SetTheme() {
	fmt.Println("执行设置主题函数, PID:", application.Get().GetPID())
	if uc.win != nil {
		uc.win.Restore()
		uc.win.Focus()
	}
}
