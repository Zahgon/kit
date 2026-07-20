package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

type TestClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	_ = "STUB: not implemented"
	return *new(TestClient)
}

func (c *testClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedTestServer()
}

type UnimplementedTestServer struct {
}

func (UnimplementedTestServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {
	_ = "STUB: not implemented"
	return
}

type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) { _ = "STUB: not implemented"; return }

func _Test_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var Test_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Test_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
