package os

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/felipecoppola/go-boilerplate/pkg/logger"

	"go.uber.org/zap"
)

// SignalListener registers a new OS signal listener.
// It will cancel the context when SIGTERM is received.
func SignalListener(logger logger.Logger, cancel context.CancelFunc) {
	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

		logger.Info("listening on OS signal requests")

		sig := <-osSignals
		logger.Warn("OS signal request received", zap.String("signal", sig.String()))
		cancel()
	}()
}
