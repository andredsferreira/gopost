package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	l.handler.ServeHTTP(w, r)
	slog.Info("http-request:",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("elapsed-time", time.Since(s).String()),
	)
}

func NewLogger(h http.Handler) *Logger {
	return &Logger{h}
}
