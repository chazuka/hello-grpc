// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculator.proto

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

type AdditionRequest struct {
	First                int32    `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               int32    `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdditionRequest) Reset()         { *m = AdditionRequest{} }
func (m *AdditionRequest) String() string { return proto.CompactTextString(m) }
func (*AdditionRequest) ProtoMessage()    {}
func (*AdditionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{0}
}

func (m *AdditionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdditionRequest.Unmarshal(m, b)
}
func (m *AdditionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdditionRequest.Marshal(b, m, deterministic)
}
func (m *AdditionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdditionRequest.Merge(m, src)
}
func (m *AdditionRequest) XXX_Size() int {
	return xxx_messageInfo_AdditionRequest.Size(m)
}
func (m *AdditionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdditionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdditionRequest proto.InternalMessageInfo

func (m *AdditionRequest) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *AdditionRequest) GetSecond() int32 {
	if m != nil {
		return m.Second
	}
	return 0
}

type AdditionResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdditionResponse) Reset()         { *m = AdditionResponse{} }
func (m *AdditionResponse) String() string { return proto.CompactTextString(m) }
func (*AdditionResponse) ProtoMessage()    {}
func (*AdditionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{1}
}

func (m *AdditionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdditionResponse.Unmarshal(m, b)
}
func (m *AdditionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdditionResponse.Marshal(b, m, deterministic)
}
func (m *AdditionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdditionResponse.Merge(m, src)
}
func (m *AdditionResponse) XXX_Size() int {
	return xxx_messageInfo_AdditionResponse.Size(m)
}
func (m *AdditionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdditionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdditionResponse proto.InternalMessageInfo

func (m *AdditionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{2}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Factor               int32    `protobuf:"varint,1,opt,name=factor,proto3" json:"factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{3}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetFactor() int32 {
	if m != nil {
		return m.Factor
	}
	return 0
}

type AverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{4}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{5}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa30b76d83302dd, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*AdditionRequest)(nil), "calculator.AdditionRequest")
	proto.RegisterType((*AdditionResponse)(nil), "calculator.AdditionResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*AverageRequest)(nil), "calculator.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "calculator.AverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
}

func init() {
	proto.RegisterFile("calculator/calculator.proto", fileDescriptor_efa30b76d83302dd)
}

var fileDescriptor_efa30b76d83302dd = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x49, 0x00, 0x33, 0x26, 0x80, 0xab, 0x21, 0x4d, 0x31, 0x42, 0x7a, 0x22, 0x82, 0x48,
	0xf4, 0xa4, 0x17, 0x83, 0xa2, 0x37, 0x8d, 0xa9, 0x37, 0x6f, 0xcb, 0x76, 0x20, 0x1b, 0x69, 0xb7,
	0xee, 0x6e, 0xd1, 0x9f, 0xf6, 0x1f, 0x8c, 0xdd, 0x96, 0xb6, 0x0a, 0xc1, 0xdb, 0xbe, 0x99, 0x37,
	0xef, 0x6d, 0xe6, 0x65, 0xa0, 0xc3, 0xe8, 0x92, 0xc5, 0x4b, 0xaa, 0x85, 0x3c, 0xcf, 0x9f, 0xa3,
	0x48, 0x0a, 0x2d, 0x08, 0xe4, 0x15, 0xf7, 0x06, 0x9a, 0x13, 0xdf, 0xe7, 0x9a, 0x8b, 0xd0, 0xc3,
	0xf7, 0x18, 0x95, 0x26, 0x47, 0x50, 0x9d, 0x73, 0xa9, 0xb4, 0x6d, 0xf5, 0xac, 0x7e, 0xd5, 0x33,
	0x80, 0xb4, 0xa1, 0xa6, 0x90, 0x89, 0xd0, 0xb7, 0x2b, 0x49, 0x39, 0x45, 0xee, 0x29, 0xb4, 0x72,
	0x01, 0x15, 0x89, 0x50, 0xe1, 0x0f, 0x57, 0xa2, 0x8a, 0x97, 0x99, 0x44, 0x8a, 0xdc, 0x2b, 0xe8,
	0x3e, 0x4b, 0x1e, 0xe0, 0x53, 0x1c, 0xcc, 0x50, 0x4e, 0x91, 0x89, 0x20, 0x12, 0xaa, 0x64, 0xde,
	0x86, 0x5a, 0x98, 0x74, 0xb3, 0x51, 0x83, 0xdc, 0x6b, 0xe8, 0x6d, 0x1f, 0xcd, 0x6d, 0xe7, 0x94,
	0x69, 0xb1, 0x9e, 0x35, 0xc8, 0xed, 0x43, 0x63, 0xb2, 0x42, 0x49, 0x17, 0xb8, 0xcb, 0x65, 0x00,
	0xcd, 0x35, 0x33, 0x15, 0xb5, 0xa1, 0x4e, 0x4d, 0x29, 0xe1, 0x5a, 0x5e, 0x06, 0xdd, 0x21, 0x90,
	0x07, 0x1e, 0xfa, 0x8f, 0xf4, 0x93, 0x07, 0x71, 0xb0, 0x4b, 0xfa, 0x0c, 0x0e, 0x4b, 0xec, 0xfc,
	0xcf, 0x9b, 0xe8, 0x17, 0x5f, 0x15, 0x38, 0xb8, 0x5b, 0xc7, 0xf4, 0x82, 0x72, 0xc5, 0x19, 0x92,
	0x7b, 0xd8, 0xcb, 0x96, 0x4d, 0x3a, 0xa3, 0x42, 0xb0, 0xbf, 0x32, 0x74, 0x8e, 0x37, 0x37, 0x53,
	0xd3, 0x0f, 0xb0, 0xb7, 0x2d, 0x93, 0x0c, 0x8a, 0x93, 0x3b, 0xd2, 0x72, 0x86, 0xff, 0x23, 0x1b,
	0xdb, 0xb1, 0x45, 0xa6, 0x50, 0x4f, 0xf7, 0x4b, 0x9c, 0xd2, 0x0f, 0x4b, 0xf1, 0x38, 0x9d, 0x8d,
	0x3d, 0xa3, 0xd2, 0xb7, 0x88, 0x07, 0xfb, 0x85, 0x55, 0x92, 0x93, 0x22, 0xfb, 0x6f, 0x22, 0x4e,
	0x77, 0x6b, 0x3f, 0x53, 0x1c, 0x5b, 0xb7, 0xad, 0xd7, 0x46, 0xe1, 0x64, 0xa2, 0xb7, 0xc5, 0xac,
	0x96, 0x1c, 0xcb, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc6, 0xf8, 0xce, 0x4b, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error) {
	out := new(AdditionResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Addition(context.Context, *AdditionRequest) (*AdditionResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	Average(CalculatorService_AverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Addition(ctx context.Context, req *AdditionRequest) (*AdditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addition not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(req *PrimeNumberDecompositionRequest, srv CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(srv CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Addition(ctx, req.(*AdditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _CalculatorService_Addition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculator.proto",
}
