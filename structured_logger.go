package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// StructuredLogger is the base struct that chi uses for the request logging.
type StructuredLogger struct {
	Logger *logrus.Logger
}

// NewStructuredLogger returns a structured logger for chi to use as a request
// logger.
func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

// NewLogEntry is called for each log entry that should be written.
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &LogEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["scheme"] = scheme
	logFields["method"] = r.Method
	logFields["uri"] = r.RequestURI

	entry.Logger = entry.Logger.WithFields(logFields)

	return entry
}

// LogEntry is a representation of a log entry.
type LogEntry struct {
	Logger logrus.FieldLogger
}

// Write implements the chi LogEntry writer.
func (l *LogEntry) Write(status, bytes int, elapsed time.Duration) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"code":    status,
		"length":  bytes,
		"resp_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Info()
}

// Write implements the chi LogEntry panic.
func (l *LogEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}
