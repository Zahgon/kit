package jsonrpc

import (
	"encoding/json"

	"github.com/go-kit/kit/endpoint"

	"context"
)

type EndpointCodec struct {
	Endpoint endpoint.Endpoint
	Decode   DecodeRequestFunc
	Encode   EncodeResponseFunc
}

type EndpointCodecMap map[string]EndpointCodec

type DecodeRequestFunc func(context.Context, json.RawMessage) (request interface{}, err error)

type EncodeResponseFunc func(context.Context, interface{}) (response json.RawMessage, err error)

type EncodeRequestFunc func(context.Context, interface{}) (request json.RawMessage, err error)

type DecodeResponseFunc func(context.Context, Response) (response interface{}, err error)
