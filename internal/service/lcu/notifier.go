package lcu

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func StartNotifier(win *application.WebviewWindow) {
	go func() {
		tickerTime := time.NewTicker(time.Second)       // 每秒
		tickerOthers := time.NewTicker(5 * time.Second) // 每 5 秒

		defer tickerTime.Stop()
		defer tickerOthers.Stop()

		for {
			select {
			case t := <-tickerTime.C:
				// 每秒发送当前时间字符串
				win.EmitEvent("time", t.Format(time.DateTime))

			case <-tickerOthers.C:
				// 每 5 秒发送 LCU 状态、端口、Token、召唤师信息
				port, token, err := GetLolClientApiInfo()
				status := err == nil
				win.EmitEvent("lcuStatus", map[string]any{"status": status})

				if status {
					win.EmitEvent("lcuCreds", map[string]any{
						"port":  port,
						"token": token,
					})

					InitCli(port, token)

					if info, err := (WailsAPI{}).GetCurrentSummoner(); err == nil {
						win.EmitEvent("summonerInfo", info)
					} else {
						fmt.Println("[Notifier] 获取召唤师信息失败:", err)
					}
				}
			}
		}
	}()
}
