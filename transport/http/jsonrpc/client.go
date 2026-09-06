package jsonrpc

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Client struct {
	client httptransport.HTTPClient

	tgt *url.URL

	method string

	enc            EncodeRequestFunc
	dec            DecodeResponseFunc
	before         []httptransport.RequestFunc
	after          []httptransport.ClientResponseFunc
	finalizer      httptransport.ClientFinalizerFunc
	requestID      RequestIDGenerator
	bufferedStream bool
}

type clientRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`
}

func NewClient(
	tgt *url.URL,
	method string,
	options ...ClientOption,
) *Client {
	_ = "STUB: not implemented"
	return nil
}

func DefaultRequestEncoder(_ context.Context, req interface{}) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

func DefaultResponseDecoder(_ context.Context, res Response) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClientOption func(*Client)

func SetClient(client httptransport.HTTPClient) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientBefore(before ...httptransport.RequestFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientAfter(after ...httptransport.ClientResponseFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientFinalizer(f httptransport.ClientFinalizerFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientRequestEncoder(enc EncodeRequestFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func ClientResponseDecoder(dec DecodeResponseFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type RequestIDGenerator interface {
	Generate() interface{}
}

func ClientRequestIDGenerator(g RequestIDGenerator) ClientOption {
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

type ClientFinalizerFunc func(ctx context.Context, err error)

type autoIncrementID struct {
	v *uint64
}

func NewAutoIncrementID(init uint64) RequestIDGenerator {
	_ = "STUB: not implemented"
	return *new(RequestIDGenerator)
}

func (i *autoIncrementID) Generate() interface{} { _ = "STUB: not implemented"; return nil }
