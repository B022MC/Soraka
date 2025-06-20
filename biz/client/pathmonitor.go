package client

import (
	"context"
	"fmt"
	"time"
)

// WatchPath periodically checks the stored client path and invokes cb only when it changes.
// Logs are printed for debug purposes.
func WatchPath(ctx context.Context, interval time.Duration, cb func(string)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		var lastPath string
		firstRun := true

		for {
			select {
			case <-ctx.Done():
				fmt.Println("[WatchPath] 上下文取消，退出监控")
				return
			case <-ticker.C:
				current := GetClientPath()
				fmt.Printf("[WatchPath] 当前路径: %s\n", current)

				if firstRun || current != lastPath {
					fmt.Printf("[WatchPath] 路径变化：%s → %s，触发回调\n", lastPath, current)
					lastPath = current
					firstRun = false
					if cb != nil {
						cb(current)
					}
				} else {
					fmt.Println("[WatchPath] 路径未变，跳过回调")
				}
			}
		}
	}()
}
