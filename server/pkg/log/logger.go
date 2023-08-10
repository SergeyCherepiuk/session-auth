package log

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

var logs chan message

func init() {
	logs = make(chan message)
}

func newLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{},
		),
	)
}

func HandleLogs() {
	logger := newLogger()
	for {
		logMessage := <-logs
		logger.LogAttrs(
			context.Background(),
			logMessage.Level,
			logMessage.Message,
			logMessage.Attrs...,
		)
	}
}