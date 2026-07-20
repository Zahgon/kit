package test

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
)

type service struct{}

func (service) Test(ctx context.Context, a string, b int64) (context.Context, string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), "", nil
}

func NewService() Service { _ = "STUB: not implemented"; return *new(Service) }

func makeTestEndpoint(svc Service) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

type serverBinding struct {
	pb.UnimplementedTestServer

	test grpctransport.Handler
}

func (b *serverBinding) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBinding(svc Service) *serverBinding { _ = "STUB: not implemented"; return nil }
