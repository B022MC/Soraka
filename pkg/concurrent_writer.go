package pkg

import (
	"bufio"
	"context"
	"time"
)

type ConcurrentWriter struct {
	w      *bufio.Writer
	jobs   chan []byte
	done   chan struct{}
	ctx    context.Context
	cancel context.CancelFunc
}

func NewConcurrentWriter(w *bufio.Writer) *ConcurrentWriter {
	ctx, cancel := context.WithCancel(context.Background())
	cw := &ConcurrentWriter{
		w:      w,
		jobs:   make(chan []byte, 1000),
		done:   make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
	cw.consume()
	return cw
}

func (w ConcurrentWriter) Write(msg []byte) (n int, err error) {
	if IsCtxDone(w.ctx) {
		return 0, nil
	}
	bts := make([]byte, len(msg))
	copy(bts, msg)
	w.jobs <- bts
	return len(msg), nil
}

func (w ConcurrentWriter) Sync() error {
	return nil
}

func (w ConcurrentWriter) consume() {
	go func() {
		flushTicker := time.NewTicker(time.Second)
		for msg := range w.jobs {
			_, _ = w.w.Write(msg)
			if len(w.jobs) == 0 {
				_ = w.w.Flush()
				continue
			}
			select {
			case <-flushTicker.C:
				_ = w.w.Flush()
			default:
			}
		}
		w.done <- struct{}{}
		flushTicker.Stop()
	}()
}

func (w ConcurrentWriter) Close(ctx context.Context) error {
	w.cancel()
	close(w.jobs)
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-w.done:
			break loop
		default:
			time.Sleep(time.Millisecond * 5)
		}
	}
	return w.w.Flush()
}
