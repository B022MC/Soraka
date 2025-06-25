package server

import (
	"Soraka/const/icon_map"
	"Soraka/internal/service/lcu"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

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
	go p.watchLCU()
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

	p.mu.Lock()
	p.LcuRP = rp
	p.LcuActive = true
	p.mu.Unlock()

	p.logger.Println("LCU 代理初始化成功")
}

func (p *Prophet) watchLCU() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	iconInited := false

	for {
		select {
		case <-ticker.C:
			port, token, err := lcu.GetLolClientApiInfo()
			if err != nil {
				p.mu.Lock()
				p.LcuActive = false
				p.mu.Unlock()
				iconInited = false
				continue
			}

			p.mu.Lock()
			needUpdate := p.LcuRP == nil || p.LcuRP.Port != port || p.LcuRP.Token != token
			p.mu.Unlock()

			if needUpdate {
				rp, err := lcu.NewRP(port, token)
				if err != nil {
					p.logger.Printf("重建 LCU 代理失败: %v", err)
					continue
				}

				p.mu.Lock()
				p.LcuRP = rp
				p.LcuActive = true
				p.mu.Unlock()

				p.logger.Printf("LCU 代理已刷新 port=%d", port)
			}

			// 在首次建立代理成功时加载 icon map
			if !iconInited && p.LcuActive {
				itemURL := fmt.Sprintf("http://127.0.0.1:8200/lcu/proxy/lol-game-data/assets/v1/items.json")
				//champURL := fmt.Sprintf("http://127.0.0.1:8200/lcu/proxy/lol-game-data/assets/v1/champions.json")
				spellURL := fmt.Sprintf("http://localhost:8200/lcu/proxy/lol-game-data/assets/v1/summoner-spells.json")

				if err := icon_map.NewIconMapDownloader(itemURL, "", spellURL); err != nil {
					p.logger.Printf("初始化图标 map 失败: %v", err)
				} else {
					p.logger.Println("图标 map 初始化完成")
					iconInited = true
				}
			}

		case <-p.ctx.Done():
			p.logger.Println("停止 LCU 状态监测")
			return
		}
	}
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
