package router

import (
	"Soraka/biz/client"
	"Soraka/biz/lcu"
	"context"
	"net/http"
	"time"
)

// Router exposes HTTP handlers for SSE and starts background watchers.
type Router struct {
	broker *SSEBroker
}

func NewRouter() *Router {
	return &Router{broker: NewSSEBroker()}
}

// Handler returns the HTTP handler for this router.
func (r *Router) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/events", r.broker)
	return mux
}

// Start begins background monitoring goroutines.
func (r *Router) Start(ctx context.Context) {
	client.WatchPath(ctx, 10*time.Second, func(p string) {
		r.broker.Broadcast("clientPath", p)
	})
	lcu.WatchLogin(ctx, 5*time.Second, func(ok bool) {
		r.broker.Broadcast("lcuStatus", ok)
		if ok {
			port, token, err := lcu.GetCredentials()
			if err == nil {
				r.broker.Broadcast("lcuCreds", map[string]string{"port": port, "token": token})
			}
		} else {
			r.broker.Broadcast("lcuCreds", map[string]string{"port": "", "token": ""})
		}
	})
}
