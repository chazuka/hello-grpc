// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greet.proto

package pkg

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type GreetResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GreetServerStreamRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetServerStreamRequest) Reset()         { *m = GreetServerStreamRequest{} }
func (m *GreetServerStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GreetServerStreamRequest) ProtoMessage()    {}
func (*GreetServerStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{3}
}

func (m *GreetServerStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetServerStreamRequest.Unmarshal(m, b)
}
func (m *GreetServerStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetServerStreamRequest.Marshal(b, m, deterministic)
}
func (m *GreetServerStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetServerStreamRequest.Merge(m, src)
}
func (m *GreetServerStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GreetServerStreamRequest.Size(m)
}
func (m *GreetServerStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetServerStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetServerStreamRequest proto.InternalMessageInfo

func (m *GreetServerStreamRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *GreetServerStreamRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GreetServerStreamResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetServerStreamResponse) Reset()         { *m = GreetServerStreamResponse{} }
func (m *GreetServerStreamResponse) String() string { return proto.CompactTextString(m) }
func (*GreetServerStreamResponse) ProtoMessage()    {}
func (*GreetServerStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{4}
}

func (m *GreetServerStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetServerStreamResponse.Unmarshal(m, b)
}
func (m *GreetServerStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetServerStreamResponse.Marshal(b, m, deterministic)
}
func (m *GreetServerStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetServerStreamResponse.Merge(m, src)
}
func (m *GreetServerStreamResponse) XXX_Size() int {
	return xxx_messageInfo_GreetServerStreamResponse.Size(m)
}
func (m *GreetServerStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetServerStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetServerStreamResponse proto.InternalMessageInfo

func (m *GreetServerStreamResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GreetClientStreamRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetClientStreamRequest) Reset()         { *m = GreetClientStreamRequest{} }
func (m *GreetClientStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GreetClientStreamRequest) ProtoMessage()    {}
func (*GreetClientStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{5}
}

func (m *GreetClientStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetClientStreamRequest.Unmarshal(m, b)
}
func (m *GreetClientStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetClientStreamRequest.Marshal(b, m, deterministic)
}
func (m *GreetClientStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetClientStreamRequest.Merge(m, src)
}
func (m *GreetClientStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GreetClientStreamRequest.Size(m)
}
func (m *GreetClientStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetClientStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetClientStreamRequest proto.InternalMessageInfo

func (m *GreetClientStreamRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type GreetClientStreamResponse struct {
	Greeting             []string `protobuf:"bytes,1,rep,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetClientStreamResponse) Reset()         { *m = GreetClientStreamResponse{} }
func (m *GreetClientStreamResponse) String() string { return proto.CompactTextString(m) }
func (*GreetClientStreamResponse) ProtoMessage()    {}
func (*GreetClientStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{6}
}

func (m *GreetClientStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetClientStreamResponse.Unmarshal(m, b)
}
func (m *GreetClientStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetClientStreamResponse.Marshal(b, m, deterministic)
}
func (m *GreetClientStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetClientStreamResponse.Merge(m, src)
}
func (m *GreetClientStreamResponse) XXX_Size() int {
	return xxx_messageInfo_GreetClientStreamResponse.Size(m)
}
func (m *GreetClientStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetClientStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetClientStreamResponse proto.InternalMessageInfo

func (m *GreetClientStreamResponse) GetGreeting() []string {
	if m != nil {
		return m.Greeting
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "greet.Person")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetServerStreamRequest)(nil), "greet.GreetServerStreamRequest")
	proto.RegisterType((*GreetServerStreamResponse)(nil), "greet.GreetServerStreamResponse")
	proto.RegisterType((*GreetClientStreamRequest)(nil), "greet.GreetClientStreamRequest")
	proto.RegisterType((*GreetClientStreamResponse)(nil), "greet.GreetClientStreamResponse")
}

func init() {
	proto.RegisterFile("greet/greet.proto", fileDescriptor_285dce36ca2d93c1)
}

var fileDescriptor_285dce36ca2d93c1 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0xc6, 0x59, 0x25, 0xa1, 0x99, 0xda, 0x43, 0x57, 0x91, 0x1a, 0x11, 0xcb, 0x82, 0x50, 0x10,
	0xaa, 0x44, 0xc4, 0xb3, 0x0f, 0xf0, 0x26, 0x92, 0x5e, 0xd4, 0x8b, 0xa4, 0x32, 0x86, 0x60, 0xb3,
	0x89, 0xbb, 0x1b, 0xff, 0x74, 0xcf, 0x92, 0xc9, 0x1a, 0x12, 0x5c, 0x41, 0x7b, 0x09, 0xcc, 0x23,
	0xbf, 0xef, 0x9b, 0xd9, 0x81, 0x71, 0xaa, 0x10, 0xcd, 0x09, 0x7d, 0xe7, 0xa5, 0x2a, 0x4c, 0xc1,
	0x3d, 0x0a, 0xc4, 0x0d, 0xf8, 0xf7, 0xa8, 0x74, 0x21, 0xf9, 0x01, 0xc0, 0x6b, 0xa6, 0xb4, 0x79,
	0x96, 0x49, 0x8e, 0x13, 0x36, 0x65, 0xb3, 0x20, 0x0e, 0x28, 0x73, 0x97, 0xe4, 0xc8, 0xf7, 0x21,
	0x58, 0x25, 0xdf, 0xd5, 0x0d, 0xaa, 0x0e, 0xea, 0x44, 0x5d, 0x14, 0xe7, 0xb0, 0x75, 0x5b, 0xe3,
	0x62, 0x7c, 0xaf, 0x50, 0x1b, 0x7e, 0x04, 0x7e, 0x49, 0x54, 0xe2, 0x0c, 0xa3, 0xd1, 0xbc, 0x91,
	0x6e, 0xa4, 0x62, 0x5b, 0x14, 0xc7, 0x30, 0xb2, 0xbf, 0xe9, 0xb2, 0x90, 0x1a, 0x79, 0x08, 0x03,
	0x6a, 0xcc, 0x64, 0x6a, 0x1d, 0xb4, 0xb1, 0x78, 0x84, 0x09, 0x35, 0x2f, 0x50, 0x7d, 0xa0, 0x5a,
	0x18, 0x85, 0x49, 0xfe, 0x3f, 0x3d, 0xbe, 0x0b, 0xbe, 0xac, 0xf2, 0x25, 0x2a, 0x1a, 0xc0, 0x8b,
	0x6d, 0x24, 0x2e, 0x60, 0xcf, 0x81, 0xfe, 0x83, 0xa7, 0x4b, 0xeb, 0xe9, 0x7a, 0x95, 0xa1, 0x34,
	0xeb, 0x78, 0x6a, 0xb5, 0xfb, 0x08, 0xa7, 0xf6, 0x66, 0x57, 0x3b, 0xfa, 0x64, 0x76, 0xe9, 0xb5,
	0xeb, 0xec, 0x05, 0x79, 0x04, 0x1e, 0xc5, 0x7c, 0xdb, 0x2a, 0x75, 0x9f, 0x24, 0xdc, 0xe9, 0x27,
	0xad, 0xc0, 0x03, 0x8c, 0x7f, 0x4c, 0xce, 0x0f, 0xbb, 0xad, 0x8e, 0x75, 0x87, 0xd3, 0xdf, 0x1b,
	0x1a, 0xee, 0x29, 0x6b, 0xc9, 0xdd, 0xb9, 0xfa, 0x64, 0xc7, 0xd2, 0xfa, 0x64, 0xd7, 0x4a, 0x66,
	0xec, 0x6a, 0xf8, 0x14, 0x34, 0xe7, 0x5c, 0xbe, 0xa5, 0x4b, 0x9f, 0xae, 0xf9, 0xec, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x61, 0xb0, 0x5c, 0x59, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetServerStream(ctx context.Context, in *GreetServerStreamRequest, opts ...grpc.CallOption) (GreetService_GreetServerStreamClient, error)
	GreetClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetClientStreamClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetServerStream(ctx context.Context, in *GreetServerStreamRequest, opts ...grpc.CallOption) (GreetService_GreetServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetServerStreamClient interface {
	Recv() (*GreetServerStreamResponse, error)
	grpc.ClientStream
}

type greetServiceGreetServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetServerStreamClient) Recv() (*GreetServerStreamResponse, error) {
	m := new(GreetServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/GreetClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetClientStreamClient{stream}
	return x, nil
}

type GreetService_GreetClientStreamClient interface {
	Send(*GreetClientStreamRequest) error
	CloseAndRecv() (*GreetClientStreamResponse, error)
	grpc.ClientStream
}

type greetServiceGreetClientStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetClientStreamClient) Send(m *GreetClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetClientStreamClient) CloseAndRecv() (*GreetClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetServerStream(*GreetServerStreamRequest, GreetService_GreetServerStreamServer) error
	GreetClientStream(GreetService_GreetClientStreamServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetServerStream(req *GreetServerStreamRequest, srv GreetService_GreetServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetServerStream not implemented")
}
func (*UnimplementedGreetServiceServer) GreetClientStream(srv GreetService_GreetClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetClientStream not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetServerStream(m, &greetServiceGreetServerStreamServer{stream})
}

type GreetService_GreetServerStreamServer interface {
	Send(*GreetServerStreamResponse) error
	grpc.ServerStream
}

type greetServiceGreetServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetServerStreamServer) Send(m *GreetServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_GreetClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetClientStream(&greetServiceGreetClientStreamServer{stream})
}

type GreetService_GreetClientStreamServer interface {
	SendAndClose(*GreetClientStreamResponse) error
	Recv() (*GreetClientStreamRequest, error)
	grpc.ServerStream
}

type greetServiceGreetClientStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetClientStreamServer) SendAndClose(m *GreetClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetClientStreamServer) Recv() (*GreetClientStreamRequest, error) {
	m := new(GreetClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetServerStream",
			Handler:       _GreetService_GreetServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetClientStream",
			Handler:       _GreetService_GreetClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greet.proto",
}
