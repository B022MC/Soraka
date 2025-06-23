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
	"io/fs"
	"log"
	"net"
)

// 嵌入前端构建产物
//
//go:embed all:frontend/dist
var assets embed.FS

// distFS is a sub filesystem rooted at the dist directory. It is used by Wails
// to serve the bundled frontend assets.
var distFS fs.FS

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
	mustCheckPortAvailable("8200") // Gin API
	mustCheckPortAvailable("9245") // Vite dev server（或 Wails dev）

	logger := log.Default()

	// Prepare embedded frontend assets. When the project hasn't been built
	// yet, this directory may be empty which will result in a missing
	// index.html warning from Wails.
	var err error
	distFS, err = fs.Sub(assets, "frontend/dist")
	if err != nil {
		logger.Printf("前端资源未找到，尝试在 'frontend/dist' 目录执行 'npm run build'")
	}
	if _, err := fs.ReadFile(distFS, "index.html"); err != nil {
		logger.Printf("未发现 index.html，需先构建前端资源: npm run build")
	}

	// 临时 app 创建主窗口
	tempApp := application.New(application.Options{})
	mainWin := tempApp.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:            "Soraka - 游戏大厅",
		Frameless:        true,
		URL:              "/",
		MinWidth:         1280,
		MinHeight:        800,
		BackgroundColour: application.NewRGBA(255, 255, 255, 90),
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Windows: application.WindowsWindow{
			Theme:                             0,
			BackdropType:                      application.Acrylic,
			HiddenOnTaskbar:                   false,
			DisableFramelessWindowDecorations: false,
		},
	})

	// 初始化业务 Router 并收集服务
	rootRouter := router.NewRootRouter(mainWin)
	var services []application.Service
	rootRouter.CollectWailsServices(&services)

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

	// 启动 LCU Proxy
	go func() {
		if err := prophet.Run(); err != nil {
			logger.Printf("LCU Proxy 停止: %v", err)
		}
	}()

	// 初始化 Wails 应用
	app := application.New(application.Options{
		Name:        "SorakaGui",
		Description: "Soraka GUI基础框架帮助开发者快速开发桌面应用",
		Services:    services,
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(distFS),
			FS:      distFS,
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

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

	logger.Println("应用启动完成，等待用户操作...")

	// 启动 Wails
	if err := app.Run(); err != nil {
		logger.Fatalf("应用退出异常: %v", err)
	}
}
