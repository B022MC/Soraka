package main

import (
	backend "Soraka/internal/app"
	"Soraka/internal/dal/logger"
	example "Soraka/internal/service/example"
	service "Soraka/internal/service/greet"
	lcuService "Soraka/internal/service/lcu"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"net"
)

// 前端构建产物
//
//go:embed all:frontend/dist
var assets embed.FS

// 托盘图标
//
//go:embed build/logo.png
var trayIcon []byte

func mustCheckPortAvailable(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("端口 %s 已被占用，请关闭其他实例后重试", port)
	}
	_ = ln.Close()
}
func main() {
	mustCheckPortAvailable("8200") // Gin API
	mustCheckPortAvailable("9245") // Vite Dev Server
	app := application.New(application.Options{
		Name:        "SorakaGui",
		Description: "Soraka GUI基础框架帮助开发者快速开发桌面应用",
		Services: []application.Service{
			application.NewService(&example.API{}),
			application.NewService(&lcuService.WailsAPI{}),
			application.NewService(&service.GreetService{}),
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
	//go func() {
	//	r := gin.Default()
	//
	//	// 初始化 LCU 端口和 token
	//	port, token, err := lcuService.GetLolClientApiInfo()
	//	if err != nil {
	//		log.Printf("初始化 LCU 代理失败: %v", err)
	//		// 这里可以考虑降级处理或延迟初始化
	//	}
	//
	//	var rp *lcuService.RP
	//	if port > 0 && token != "" {
	//		rp, err = lcuService.NewRP(port, token)
	//		if err != nil {
	//			log.Printf("创建 LCU 反向代理失败: %v", err)
	//		}
	//	}
	//
	//	api := &Api{
	//		p: &Prophet{
	//			LcuActive: true,
	//			lcuRP:     rp,
	//		},
	//	}
	//
	//	RegisterRoutes(r, api)
	//
	//	if err := r.Run(":8200"); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	// 初始化 logger
	logger.Init()

	prophet := backend.NewProphet()
	go func() {
		if err := prophet.Run(); err != nil {
			log.Fatal("启动 Prophet 失败:", err)
		}
	}()

	lcuService.StartNotifier(mainWin)
	// 监听前端事件
	app.OnEvent("myevent", func(e *application.CustomEvent) {
		fmt.Println("main监听事件：", e)
	})

	// 启动应用
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
