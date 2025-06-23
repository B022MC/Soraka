package greet

// WailsService exposes only methods intended for the frontend.
type WailsService struct {
	svc *GreetService
}

// NewWailsService wraps a GreetService for use with Wails.
func NewWailsService(svc *GreetService) *WailsService {
	return &WailsService{svc: svc}
}

// Name returns the service name.
func (w *WailsService) Name() string {
	return w.svc.Name()
}

// SetTheme forwards the call to the underlying service.
func (w *WailsService) SetTheme() {
	w.svc.SetTheme()
}
