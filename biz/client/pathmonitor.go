package client

import (
	"context"
	"time"
)

// WatchPath periodically checks the stored client path and invokes cb every interval.
func WatchPath(ctx context.Context, interval time.Duration, cb func(string)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		path := GetClientPath()
		if cb != nil {
			cb(path)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				path = GetClientPath()
				if cb != nil {
					cb(path)
				}
			}
		}
	}()
}
