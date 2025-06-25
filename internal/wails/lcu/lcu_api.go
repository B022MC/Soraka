package lcu

import (
	lcuReq "Soraka/internal/dal/req/lcu"
	lcuResp "Soraka/internal/dal/resp/lcu"
	lcuSvc "Soraka/internal/service/lcu"
	"log"
)

// LcuApiWails exposes only methods intended for the frontend.
type LcuApiWails struct {
	svc *lcuSvc.LcuApiService
	log *log.Logger
}

// NewWailsService wraps a LcuApiService for use with Wails.
func NewLcuApiWails(svc *lcuSvc.LcuApiService, log *log.Logger) *LcuApiWails {
	return &LcuApiWails{
		svc: svc,
		log: log,
	}
}

// GetClientPath forwards the call to the underlying service.
func (w *LcuApiWails) GetClientPath() (string, error) {
	return w.svc.GetClientPath()
}

func (w *LcuApiWails) StartClient() error {
	return w.svc.StartClient()
}
func (w *LcuApiWails) ListRecentMatches(limit int) []*lcuResp.MatchBrief {
	req := &lcuReq.MatchListReq{
		Limit: limit,
	}
	list, err := w.svc.ListRecentMatches(req)
	if err != nil {
		log.Fatalf("获取生涯失败：%v", err)
		return nil
	}
	return list
}
