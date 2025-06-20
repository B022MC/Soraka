package service

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct{}

var mainWindow *application.WebviewWindow

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

// 设置主题
func (g *GreetService) SetTheme() {
	// application.OpenFileDialog().PromptForMultipleSelection()
	fmt.Println("执行设置主题函数", application.Get().GetPID())
	if mainWindow != nil {
		mainWindow.Restore()
		mainWindow.Focus()
	}
}
