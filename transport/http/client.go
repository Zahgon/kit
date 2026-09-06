package http

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	client         HTTPClient
	req            CreateRequestFunc
	dec            DecodeResponseFunc
	before         []RequestFunc
	after          []ClientResponseFunc
	finalizer      []ClientFinalizerFunc
	bufferedStream bool
}

func NewClient(method string, tgt *url.URL, enc EncodeRequestFunc, dec DecodeResponseFunc, options ...ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

func NewExplicitClient(req CreateRequestFunc, dec DecodeResponseFunc, options ...ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

type ClientOption func(*Client)

func SetClient(client HTTPClient) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientBefore(before ...RequestFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientAfter(after ...ClientResponseFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientFinalizer(f ...ClientFinalizerFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func BufferedStream(buffered bool) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func (c Client) Endpoint() endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

type bodyWithCancel struct {
	io.ReadCloser

	cancel context.CancelFunc
}

func (bwc bodyWithCancel) Close() error { _ = "STUB: not implemented"; return nil }

type ClientFinalizerFunc func(ctx context.Context, err error)

func EncodeJSONRequest(c context.Context, r *http.Request, request interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func EncodeXMLRequest(c context.Context, r *http.Request, request interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func makeCreateRequestFunc(method string, target *url.URL, enc EncodeRequestFunc) CreateRequestFunc {
	_ = "STUB: not implemented"
	return *new(CreateRequestFunc)
}
