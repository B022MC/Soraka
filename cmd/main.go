package main

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
       "github.com/wailsapp/wails/v3/pkg/application"

"Soraka"
example "Soraka/service/example"
lcuService "Soraka/service/lcu"
)

// 前端构建产物
//
//go:embed all:../frontend/dist
var assets embed.FS

// 托盘图标
//
//go:embed logo.png
var trayIcon []byte

func main() {
	app := application.New(application.Options{
		Name:        "SorakaGui",
		Description: "Soraka GUI基础框架帮助开发者快速开发桌面应用",
               Services:    []application.Service{
                       &example.API{},
                       &lcuService.WailsAPI{},
               },

		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// 创建主窗口
	mainWin := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Soraka - 游戏大厅",
		Frameless: true, // 无边框，前端实现拖动
		URL:       "/",
		//BackgroundColour: application.NewRGB(255, 255, 255),
		MinWidth:         1280,                                   //最小宽度
		MinHeight:        800,                                    //最小高度
		BackgroundColour: application.NewRGBA(255, 255, 255, 90), // 透明度较低以显毛玻璃

		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Windows: application.WindowsWindow{
			Theme:                             0,
			BackdropType:                      application.Acrylic,
			HiddenOnTaskbar:                   false,
			DisableFramelessWindowDecorations: false, // 允许拖动阴影
		},
	})

	// === 托盘图标设置 ===
	tray := app.NewSystemTray()
	tray.SetLabel("Soraka")
	tray.SetIcon(trayIcon)
	tray.SetDarkModeIcon(trayIcon)

	// 双击左键：显示窗口
	tray.OnDoubleClick(func() {
		mainWin.Show()
		mainWin.Focus()
	})

	// start API server for frontend calls
	go func() {
		r := gin.Default()
		api := &Soraka.Api{}
		Soraka.RegisterRoutes(r, api)
		if err := r.Run(":8200"); err != nil {
			log.Fatal(err)
		}
	}()

	// 每秒发送时间事件
	go func() {
		fmt.Println("窗口id：", mainWin.ID())
		for {
			app.EmitEvent("time", time.Now().Format(time.DateTime))
			time.Sleep(time.Second)
		}
	}()

	// 监听前端事件
	app.OnEvent("myevent", func(e *application.CustomEvent) {
		fmt.Println("main监听事件：", e)
	})

	// 启动应用
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
