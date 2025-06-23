package greet

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetUseCase struct {
	mainWindow *application.WebviewWindow
}

func NewGreetUseCase(win *application.WebviewWindow) *GreetUseCase {
	return &GreetUseCase{
		mainWindow: win,
	}
}

// 核心业务逻辑：返回问候语
func (uc *GreetUseCase) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// 核心业务逻辑：设置主题/窗口操作
func (uc *GreetUseCase) SetTheme() {
	fmt.Println("执行设置主题函数, PID:", application.Get().GetPID())
	if uc.mainWindow != nil {
		uc.mainWindow.Restore()
		uc.mainWindow.Focus()
	}
}
