package greet

import greetSvc "Soraka/internal/service/greet"

// GreetSvcWails exposes only methods intended for the frontend.
type GreetSvcWails struct {
	svc *greetSvc.GreetService
}

// NewWailsService wraps a GreetService for use with Wails.
func NewGreetSvcWails(svc *greetSvc.GreetService) *GreetSvcWails {
	return &GreetSvcWails{svc: svc}
}

// Name returns the service name.
func (w *GreetSvcWails) Name() string {
	return w.svc.Name()
}

// SetTheme forwards the call to the underlying service.
func (w *GreetSvcWails) SetTheme() {
	w.svc.SetTheme()
}
