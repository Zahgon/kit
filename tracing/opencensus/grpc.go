package opencensus

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

const propagationKey = "grpc-trace-bin"

func GRPCClientTrace(options ...TracerOption) kitgrpc.ClientOption {
	_ = "STUB: not implemented"
	return *new(kitgrpc.ClientOption)
}

func GRPCServerTrace(options ...TracerOption) kitgrpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(kitgrpc.ServerOption)
}
