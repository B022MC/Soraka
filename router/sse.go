package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type message struct {
	event string
	data  string
}

// SSEBroker manages client connections for Server-Sent Events.
type SSEBroker struct {
	mu      sync.Mutex
	clients map[chan message]struct{}
}

// NewSSEBroker creates a broker.
func NewSSEBroker() *SSEBroker {
	return &SSEBroker{clients: make(map[chan message]struct{})}
}

// ServeHTTP handles incoming SSE connections at /events.
func (b *SSEBroker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan message, 10)

	b.mu.Lock()
	b.clients[ch] = struct{}{}
	b.mu.Unlock()

	defer func() {
		b.mu.Lock()
		delete(b.clients, ch)
		b.mu.Unlock()
	}()

	// send an initial comment so proxies establish the stream immediately
	fmt.Fprint(w, ": connected\n\n")
	flusher.Flush()

	notify := r.Context().Done()
	keepAlive := time.NewTicker(10 * time.Second)
	defer keepAlive.Stop()
	for {
		select {
		case <-notify:
			return
		case <-keepAlive.C:
			fmt.Fprint(w, ": heartbeat\n\n")
			flusher.Flush()
		case msg := <-ch:
			fmt.Fprintf(w, "event: %s\n", msg.event)
			fmt.Fprintf(w, "data: %s\n\n", msg.data)
			flusher.Flush()
		}
	}
}

// Broadcast sends an event with JSON data to all clients.
func (b *SSEBroker) Broadcast(event string, v interface{}) {
	data, _ := json.Marshal(v)
	msg := message{event: event, data: string(data)}
	b.mu.Lock()
	for ch := range b.clients {
		select {
		case ch <- msg:
		default:
		}
	}
	b.mu.Unlock()
}

// StartHeartbeat periodically sends a ping event to all clients to keep connections alive.
func (b *SSEBroker) StartHeartbeat(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			b.Broadcast("ping", "pong")
		}
	}()
}
