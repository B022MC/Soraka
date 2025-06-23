package app

import (
	"Soraka/internal/dal/lcu/models"
	"Soraka/internal/dal/logger"
	"Soraka/internal/global"
	"Soraka/internal/service/lcu"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	bdkgin "Soraka/pkg/gin"
	bdkmid "Soraka/pkg/gin/middleware"
)

type (
	GameState string
	Prophet   struct {
		ctx          context.Context
		opts         *options
		httpSrv      *http.Server
		lcuPort      int
		lcuToken     string
		LcuActive    bool
		CurrSummoner *models.SummonerProfileData
		cancel       func()
		api          *Api
		mu           *sync.Mutex
		lcuRP        *lcu.RP
	}
	options struct {
		debug       bool
		enablePprof bool
		httpAddr    string
	}
)

// gameState
const (
	GameStateNone GameState = "none"
)

var (
	defaultOpts = &options{
		debug:       false,
		enablePprof: true,
		httpAddr:    "127.0.0.1:8200",
	}
)

func NewProphet(opts ...ApplyOption) *Prophet {
	ctx, cancel := context.WithCancel(context.Background())
	p := &Prophet{
		ctx:    ctx,
		cancel: cancel,
		mu:     &sync.Mutex{},
		opts:   defaultOpts,
	}
	if global.IsDevMode() {
		opts = append(opts, WithDebug())
	} else {
		opts = append(opts, WithProd())
	}
	p.api = &Api{p: p}
	for _, fn := range opts {
		fn(p.opts)
	}
	return p
}
func (p *Prophet) Run() error {
	p.initGin()
	return p.notifyQuit()
}
func (p *Prophet) Stop() error {
	if p.cancel != nil {
		p.cancel()
	}
	// stop all task
	return nil
}

func (p *Prophet) notifyQuit() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	g, c := errgroup.WithContext(p.ctx)
	// http
	g.Go(func() error {
		err := p.httpSrv.ListenAndServe()
		if err != nil || !errors.Is(err, http.ErrServerClosed) {
			_ = p.Stop()
			return err
		}
		return nil
	})
	// http-shutdown
	g.Go(func() error {
		<-c.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		return p.httpSrv.Shutdown(ctx)
	})
	// wait quit
	g.Go(func() error {
		for {
			select {
			case <-p.ctx.Done():
				return p.ctx.Err()
			case <-interrupt:
				_ = p.Stop()
			}
		}
	})
	err := g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
func (p *Prophet) initGin() {
	if p.opts.debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := bdkgin.NewGin()
	if p.opts.debug {
		engine.Use(gin.LoggerWithFormatter(bdkgin.LogFormatter))
	}
	if p.opts.enablePprof {
		pprof.RouteRegister(engine.Group(""))
	}
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowWebSockets:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Use(bdkmid.RecoveryWithLogFn(logger.Error))

	port, token, err := lcu.GetLolClientApiInfo()
	if err != nil {
		logger.Warn("初始化 LCU 代理失败", err)
	} else {
		rp, err := lcu.NewRP(port, token)
		if err != nil {
			logger.Warn("创建 LCU 反向代理失败", err)
		} else {
			p.lcuRP = rp
			p.LcuActive = true
			logger.Info("LCU 代理初始化成功")
		}
	}

	p.api = &Api{p: p}
	RegisterRoutes(engine, p.api)

	srv := &http.Server{
		Addr:    p.opts.httpAddr,
		Handler: engine,
	}
	p.httpSrv = srv

	logger.Info(fmt.Sprintf("Gin 服务已初始化，监听地址: %s", p.opts.httpAddr))
}
