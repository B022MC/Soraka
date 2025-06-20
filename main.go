package main

import (
	"Soraka/service"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "SorakaGui",
		Description: "Soraka GUI基础框架帮助开发者快速开发桌面应用",
		Services: []application.Service{
			application.NewService(&service.GreetService{}),
			application.NewService(&service.MessageService{}),
			application.NewService(&service.OpenWindow{}),
			application.NewService(&service.HttpService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	mainWind := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Soraka - 游戏大厅",
		// Frameless: true,//移除窗口的窗框(即边框、头部标题和操作按钮)，用于自定义窗口头部
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(255, 255, 255),
		Windows:          application.WindowsWindow{Theme: 0}, //这里是设框架主题，0=跟随系统，1=Dark(黑色)，2=Light(浅色)
		URL:              "/",
	})
	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		fmt.Println("窗口id：", mainWind.ID())
		for {
			now := time.Now().Format(time.DateTime)
			app.EmitEvent("time", now)
			time.Sleep(time.Second)
		}
	}()
	//处理监听前端自定义事件
	app.OnEvent("myevent", func(e *application.CustomEvent) {
		// Access event information
		fmt.Println("main监听事件：", e)
		return
	})

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}

}
