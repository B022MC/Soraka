package lcu

// WailsService exposes only methods intended for the frontend.
type WailsService struct {
	svc *LcuApiService
}

// NewWailsService wraps a LcuApiService for use with Wails.
func NewWailsService(svc *LcuApiService) *WailsService {
	return &WailsService{svc: svc}
}

// Name returns the service name.
func (w *WailsService) Name() string {
	return w.svc.Name()
}

// GetClientPath forwards the call to the underlying service.
func (w *WailsService) GetClientPath() (string, error) {
	return w.svc.GetClientPath()
}
