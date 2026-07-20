package http

import (
	"net/http"
)

type interceptingWriter struct {
	http.ResponseWriter
	code    int
	written int64
}

func (w *interceptingWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (w *interceptingWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *interceptingWriter) reimplementInterfaces() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}
