package service

import "github.com/wailsapp/wails/v3/pkg/application"

type OpenWindow struct{}

func (g *OpenWindow) Open(url string) string {
	app := application.New(application.Options{
		Name: "打开新窗口",
	})
	window2 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "新打开的演示窗口",
		Width:  1300,
		Height: 800,
	})
	// Load an external URL
	window2.SetURL(url)
	return "窗口打开成功"
}
