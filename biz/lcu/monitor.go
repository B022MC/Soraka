package lcu

import (
	"context"
	"time"
)

// WatchLogin periodically checks login status and credentials and notifies via callback.
// The callback is invoked on every tick to act as a heartbeat.
func WatchLogin(ctx context.Context, interval time.Duration, cb func(bool, string, string)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		// send initial status immediately
		ok, port, token := CheckLogin()
		if cb != nil {
			cb(ok, port, token)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ok, port, token = CheckLogin()
				if cb != nil {
					cb(ok, port, token)
				}
			}
		}
	}()
}
