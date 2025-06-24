package server

import (
	"Soraka/internal/biz/lcu"
	"Soraka/pkg"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Notifier struct {
	win *application.WebviewWindow
}

func NewNotifier(win *application.WebviewWindow) *Notifier {
	return &Notifier{win: win}
}

// Start 启动定时任务推送 LCU 状态和时间
func (n *Notifier) Start() {
	go func() {
		tickerTime := time.NewTicker(1 * time.Second)
		tickerOthers := time.NewTicker(5 * time.Second)
		defer tickerTime.Stop()
		defer tickerOthers.Stop()

		for {
			select {
			case t := <-tickerTime.C:
				n.win.EmitEvent("time", t.Format(time.DateTime))

			case <-tickerOthers.C:
				n.pushLCUInfo()
			}
		}
	}()
}

// pushLCUInfo 检测 LCU 状态并推送信息
func (n *Notifier) pushLCUInfo() {
	clientInfoUC := lcu.NewClientInfoUseCase()
	port, token, err := clientInfoUC.GetLolClientApiInfo()

	status := err == nil
	n.win.EmitEvent("lcuStatus", map[string]any{"status": status})

	if !status {
		fmt.Println("[Notifier] 未找到 LOL 客户端进程:", err)
		return
	}

	n.win.EmitEvent("lcuCreds", map[string]any{
		"port":  port,
		"token": token,
	})

	client := pkg.NewClient(port, token)
	logger := log.Default()
	apiUC := lcu.NewLcuApiUseCase(logger, client)

	info, err := apiUC.GetCurrSummoner()
	if err != nil {
		fmt.Println("[Notifier] 获取召唤师信息失败:", err)
		return
	}

	n.win.EmitEvent("summonerInfo", info)
}
