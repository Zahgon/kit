package zipkin

import (
	zipkin "github.com/openzipkin/zipkin-go"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

func GRPCClientTrace(tracer *zipkin.Tracer, options ...TracerOption) kitgrpc.ClientOption {
	_ = "STUB: not implemented"
	return *new(kitgrpc.ClientOption)
}

func GRPCServerTrace(tracer *zipkin.Tracer, options ...TracerOption) kitgrpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(kitgrpc.ServerOption)
}
