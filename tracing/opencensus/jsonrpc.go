package opencensus

import (
	jsonrpc "github.com/go-kit/kit/transport/http/jsonrpc"
)

func JSONRPCClientTrace(options ...TracerOption) jsonrpc.ClientOption {
	_ = "STUB: not implemented"
	return *new(jsonrpc.ClientOption)
}

func JSONRPCServerTrace(options ...TracerOption) jsonrpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(jsonrpc.ServerOption)
}
