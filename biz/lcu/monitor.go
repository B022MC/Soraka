package lcu

import (
	"context"
	"time"
)

// WatchLogin periodically checks login status and notifies via callback when it changes.
func WatchLogin(ctx context.Context, interval time.Duration, cb func(bool)) {
        go func() {
                ticker := time.NewTicker(interval)
                defer ticker.Stop()

                // send initial status
                prev := CheckLogin()
                if cb != nil {
                        cb(prev)
                }

                for {
                        select {
                        case <-ctx.Done():
                                return
                        case <-ticker.C:
                                curr := CheckLogin()
                                if curr != prev {
                                        prev = curr
                                        if cb != nil {
                                                cb(curr)
                                        }
                                }
                        }
                }
        }()
}
