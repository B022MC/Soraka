package main

import (
	"Soraka/internal/proxy"
	"Soraka/internal/router"
	"Soraka/internal/server"
	bdkgin "Soraka/pkg/gin"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"net"
)

// 嵌入前端构建产物
//
//go:embed all:frontend/dist
var assets embed.FS

// 嵌入托盘图标
//
//go:embed build/logo.png
var trayIcon []byte

// 检查端口占用
func mustCheckPortAvailable(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("端口 %s 已被占用，请关闭其他实例后重试", port)
	}
	_ = ln.Close()
}

func main() {
	//if !utils.IsRunAsAdmin() {
	//	err := utils.RelaunchAsAdmin()
	//	if err != nil {
	//		fmt.Printf("以管理员身份重新启动失败: %v\n", err)
	//		os.Exit(1)
	//	}
	//	// 当前进程退出，等待管理员权限的进程继续
	//	os.Exit(0)
	//}
	mustCheckPortAvailable("8200") // Gin API
	mustCheckPortAvailable("9245") //
	// 初始化业务 Router 并收集服务
	rootRouter := router.NewRootRouter()
	var services []application.Service
	rootRouter.CollectWailsServices(&services)

	app := application.New(application.Options{
		Name:        "SorakaGui",
		Description: "Soraka GUI基础框架帮助开发者快速开发桌面应用",
		Services:    services,

		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	logger := log.Default()

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
	// 注入 window
	rootRouter.SetMainWin(mainWin)
	// 初始化 Gin 引擎
	engine := bdkgin.NewGin()
	engine.Use(gin.Logger(), gin.Recovery())
	rootRouter.InitRouter(&engine.RouterGroup)

	// 初始化 LCU Proxy
	prophet := server.NewProphet(logger)
	proxy.RegisterLcuProxy(engine, prophet)

	// 初始化 HTTP Server
	httpServer := server.NewHTTPServer("127.0.0.1:8200", true, engine, logger)

	// 启动 Gin 服务
	go func() {
		if err := httpServer.StartWithSignal(); err != nil {
			logger.Printf("HTTP 服务启动失败: %v", err)
		}
	}()
	//// 启动 LCU Proxy
	go func() {
		if err := prophet.Run(); err != nil {
			logger.Printf("LCU Proxy 停止: %v", err)
		}
	}()
	// 启动 Notifier 定时推送任务
	notifier := server.NewNotifier(mainWin)
	notifier.Start()

	// 托盘设置
	tray := app.NewSystemTray()
	tray.SetLabel("Soraka")
	tray.SetIcon(trayIcon)
	tray.SetDarkModeIcon(trayIcon)
	tray.OnDoubleClick(func() {
		mainWin.Show()
		mainWin.Focus()
	})

	// 前端事件监听示例
	app.OnEvent("myevent", func(e *application.CustomEvent) {
		fmt.Println("main监听事件：", e)
	})

	// 启动 Wails
	if err := app.Run(); err != nil {
		logger.Fatalf("应用退出异常: %v", err)
	}
}
