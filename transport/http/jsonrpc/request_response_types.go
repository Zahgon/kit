package jsonrpc

import (
	"context"
	"encoding/json"
	"net/http"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      *RequestID      `json:"id"`
}

type RequestID struct {
	intValue    int
	intError    error
	floatValue  float32
	floatError  error
	stringValue string
	stringError error
}

type RequestFunc func(context.Context, *http.Request, Request) context.Context

func (id *RequestID) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (id *RequestID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *RequestID) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (id *RequestID) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (id *RequestID) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

type Response struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
	ID      *RequestID      `json:"id"`
}

const (
	Version string = "2.0"

	ContentType string = "application/json; charset=utf-8"
)

type contextKey int

const (
	ContextKeyRequestMethod contextKey = iota
)
