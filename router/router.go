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
	mux.Handle("/events", withCORS(r.broker))
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
	r.broker.StartHeartbeat(30 * time.Second)
}
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 添加 CORS 相关响应头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求则直接返回
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 正常请求继续
		next.ServeHTTP(w, req)
	})
}
