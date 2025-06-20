package router

import (
        "Soraka/biz/client"
        "Soraka/biz/lcu"
        "context"
        "time"

        "github.com/wailsapp/wails/v3/pkg/application"
)

// Router starts background watchers and emits events to the frontend via the Wails app.
type Router struct {
        app *application.App
}

// NewRouter creates a new Router bound to the given app.
func NewRouter(app *application.App) *Router {
        return &Router{app: app}
}

// Start begins background monitoring goroutines.
func (r *Router) Start(ctx context.Context) {
        client.WatchPath(ctx, 10*time.Second, func(p string) {
                r.app.EmitEvent("clientPath", p)
        })
        lcu.WatchLogin(ctx, 5*time.Second, func(ok bool) {
                r.app.EmitEvent("lcuStatus", ok)
                if ok {
                        port, token, err := lcu.GetCredentials()
                        if err == nil {
                                r.app.EmitEvent("lcuCreds", map[string]string{"port": port, "token": token})
                        }
                } else {
                        r.app.EmitEvent("lcuCreds", map[string]string{"port": "", "token": ""})
                }
        })
}

