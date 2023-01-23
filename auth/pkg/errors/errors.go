package errors

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

func SignalHandler(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, unix.SIGTERM, unix.SIGINT)

	select {
	case sig := <-quit:
		return fmt.Errorf("%s", sig)
	case <-ctx.Done():
		return nil
	}
}
