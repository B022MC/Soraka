package lcu

import lcuSvc "Soraka/internal/service/lcu"

// LcuApiWails exposes only methods intended for the frontend.
type LcuApiWails struct {
	svc *lcuSvc.LcuApiService
}

// NewWailsService wraps a LcuApiService for use with Wails.
func NewLcuApiWails(svc *lcuSvc.LcuApiService) *LcuApiWails {
	return &LcuApiWails{svc: svc}
}

// GetClientPath forwards the call to the underlying service.
func (w *LcuApiWails) GetClientPath() (string, error) {
	return w.svc.GetClientPath()
}

func (w *LcuApiWails) StartClient() error {
	return w.svc.StartClient()
}
