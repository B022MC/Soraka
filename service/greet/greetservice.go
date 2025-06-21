package service

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct{}

var mainWindow *application.WebviewWindow

// 实现 application.Service 接口
func (GreetService) Name() string {
	return "GreetService"
}

// 可被前端调用的方法
func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

// 设置主题（用于测试窗口操作）
func (g *GreetService) SetTheme() {
	fmt.Println("执行设置主题函数, PID:", application.Get().GetPID())

	if mainWindow != nil {
		mainWindow.Restore()
		mainWindow.Focus()
	}
}

// 可选：设置窗口引用
func SetMainWindow(win *application.WebviewWindow) {
	mainWindow = win
}
