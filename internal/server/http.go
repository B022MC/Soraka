package server

import (
	"Soraka/internal/service/lcu"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServer struct {
	addr   string
	srv    *http.Server
	debug  bool
	rp     *lcu.RP
	logger *log.Logger
}

func NewHTTPServer(addr string, debug bool, handler http.Handler, logger *log.Logger) *HTTPServer {
	return &HTTPServer{
		addr: addr,
		srv: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		debug:  debug,
		logger: logger,
	}
}

//func NewHTTPServer(addr string, debug bool, enablePprof bool, rootRouter *router.RootRouter) *HTTPServer {
//	if debug {
//		gin.SetMode(gin.DebugMode)
//	} else {
//		gin.SetMode(gin.ReleaseMode)
//	}
//
//	engine := bdkgin.NewGin()
//	if debug {
//		engine.Use(gin.LoggerWithFormatter(bdkgin.LogFormatter))
//	}
//	if enablePprof {
//		pprof.RouteRegister(engine.Group(""))
//	}
//	engine.Use(cors.New(cors.Config{
//		AllowOrigins:     []string{"*"},
//		AllowMethods:     []string{"*"},
//		AllowHeaders:     []string{"*"},
//		ExposeHeaders:    []string{"*"},
//		AllowWebSockets:  true,
//		AllowCredentials: true,
//		MaxAge:           12 * time.Hour,
//	}))
//	engine.Use(bdkmid.RecoveryWithLogFn(logger.Error))
//
//	// 调用路由初始化
//	rootRouter.InitRouter(&engine.RouterGroup)
//
//	return &HTTPServer{
//		addr: addr,
//		srv: &http.Server{
//			Addr:    addr,
//			Handler: engine,
//		},
//		debug: debug,
//	}
//}

func (h *HTTPServer) StartWithSignal() error {
	// 信号监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// 启动 HTTP
	go func() {
		h.logger.Printf("Gin 服务已初始化，监听地址: %s", h.addr)
		if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.logger.Printf("HTTP 服务异常退出: %v", err)
		}
	}()

	// 等待退出信号
	<-quit
	h.logger.Println("收到退出信号，开始关闭 HTTP 服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := h.srv.Shutdown(ctx); err != nil {
		h.logger.Printf("HTTP 服务关闭失败: %v", err)
		return err
	}

	h.logger.Println("HTTP 服务已优雅退出")
	return nil
}
