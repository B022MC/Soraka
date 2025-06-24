package server

import (
	"Soraka/internal/service/lcu"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type Prophet struct {
	ctx       context.Context
	cancel    func()
	mu        *sync.Mutex
	LcuRP     *lcu.RP
	LcuActive bool
	logger    *log.Logger
}

func NewProphet(logger *log.Logger) *Prophet {
	ctx, cancel := context.WithCancel(context.Background())
	return &Prophet{
		ctx:    ctx,
		cancel: cancel,
		mu:     &sync.Mutex{},
		logger: logger,
	}
}

func (p *Prophet) Run() error {
	p.initLCU()
	return p.notifyQuit()
}

func (p *Prophet) initLCU() {
	port, token, err := lcu.GetLolClientApiInfo()
	if err != nil {
		p.logger.Printf("初始化 LCU 代理失败: %v", err)
		return
	}
	rp, err := lcu.NewRP(port, token)
	if err != nil {
		p.logger.Printf("创建 LCU 反向代理失败: %v", err)
		return
	}
	p.LcuRP = rp
	p.LcuActive = true
	p.logger.Println("LCU 代理初始化成功")
}

func (p *Prophet) notifyQuit() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	g, ctx := errgroup.WithContext(p.ctx)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-interrupt:
				return p.Stop()
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (p *Prophet) Stop() error {
	if p.cancel != nil {
		p.cancel()
	}
	p.logger.Println("服务已停止")
	return nil
}
