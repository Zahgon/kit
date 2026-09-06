package pb

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64  `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *TestRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *TestRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*TestRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*TestRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *TestRequest) GetA() string { _ = "STUB: not implemented"; return "" }

func (x *TestRequest) GetB() int64 { _ = "STUB: not implemented"; return 0 }

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V string `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *TestResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *TestResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*TestResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*TestResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *TestResponse) GetV() string { _ = "STUB: not implemented"; return "" }

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x29, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x1c, 0x0a, 0x0c, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x76,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x32, 0x33, 0x0a, 0x04, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*TestRequest)(nil),
	(*TestResponse)(nil),
}
var file_test_proto_depIdxs = []int32{
	0,
	1,
	1,
	0,
	0,
	0,
	0,
}

func init()                 { file_test_proto_init() }
func file_test_proto_init() { _ = "STUB: not implemented"; return }
