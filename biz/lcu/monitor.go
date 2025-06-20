package lcu

import (
	"context"
	"fmt"
	"time"
)

// WatchLogin periodically checks login status and credentials and notifies via callback.
// The callback is only invoked when there is a change, or on the first check.
func WatchLogin(ctx context.Context, interval time.Duration, cb func(bool, string, string)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		var lastOK bool
		var lastPort, lastToken string
		firstRun := true

		for {
			select {
			case <-ctx.Done():
				fmt.Println("[WatchLogin] 上下文取消，退出监控循环")
				return
			case <-ticker.C:
				ok, port, token := CheckLogin()

				if firstRun || ok != lastOK || port != lastPort || token != lastToken {
					fmt.Printf("[WatchLogin] 状态变化或首次 -> ok: %v, port: %s, token: %s\n", ok, port, token)
					lastOK = ok
					lastPort = port
					lastToken = token
					firstRun = false

					if cb != nil {
						fmt.Println("[WatchLogin] 触发回调")
						cb(ok, port, token)
					}
				} else {
					fmt.Println("[WatchLogin] 状态未变，跳过回调")
				}
			}
		}
	}()
}
