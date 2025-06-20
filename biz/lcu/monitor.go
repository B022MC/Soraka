package lcu

import (
	"context"
	"time"
)

// WatchLogin periodically checks login status and notifies via callback.
// The callback is invoked on every tick to act as a heartbeat.
func WatchLogin(ctx context.Context, interval time.Duration, cb func(bool)) {
       go func() {
               ticker := time.NewTicker(interval)
               defer ticker.Stop()

               // send initial status immediately
               status := CheckLogin()
               if cb != nil {
                       cb(status)
               }

               for {
                       select {
                       case <-ctx.Done():
                               return
                       case <-ticker.C:
                               status = CheckLogin()
                               if cb != nil {
                                       cb(status)
                               }
                       }
               }
       }()
}
