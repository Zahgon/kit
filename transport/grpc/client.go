package grpc

import (
	"context"
	"reflect"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
)

type Client struct {
	client      *grpc.ClientConn
	serviceName string
	method      string
	enc         EncodeRequestFunc
	dec         DecodeResponseFunc
	grpcReply   reflect.Type
	before      []ClientRequestFunc
	after       []ClientResponseFunc
	finalizer   []ClientFinalizerFunc
}

func NewClient(
	cc *grpc.ClientConn,
	serviceName string,
	method string,
	enc EncodeRequestFunc,
	dec DecodeResponseFunc,
	grpcReply interface{},
	options ...ClientOption,
) *Client {
	_ = "STUB: not implemented"
	return nil
}

type ClientOption func(*Client)

func ClientBefore(before ...ClientRequestFunc) ClientOption {
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

func (c Client) Endpoint() endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

type ClientFinalizerFunc func(ctx context.Context, err error)
