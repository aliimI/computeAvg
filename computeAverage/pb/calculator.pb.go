package pb

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

type ComputeAverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a57f8e0def5b4cf, []int{0}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type ComputeAverageRequest struct {
	Val                  int32    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a57f8e0def5b4cf, []int{1}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*ComputeAverageResponse)(nil), "computeAverage.ComputeAverageResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "computeAverage.ComputeAverageRequest")
}

func init() { proto.RegisterFile("computeAverage.proto", fileDescriptor_7a57f8e0def5b4cf) }

var fileDescriptor_7a57f8e0def5b4cf = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0x2d,
	0x28, 0x2d, 0x49, 0x75, 0x2c, 0x4b, 0x2d, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x55, 0x32, 0xe2, 0x12, 0x73, 0x46, 0x11, 0x09, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0x84, 0x08, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x30,
	0x06, 0xc1, 0xb8, 0x4a, 0x9a, 0x5c, 0xa2, 0xe8, 0x7a, 0x0a, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04,
	0xb8, 0x98, 0xcb, 0x12, 0x73, 0xc0, 0xca, 0x59, 0x83, 0x40, 0x4c, 0xa3, 0x0a, 0x74, 0xa5, 0xc1,
	0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xf1, 0x5c, 0x5c, 0x70, 0x89, 0x74, 0x21, 0x55, 0x3d,
	0x34, 0xc7, 0x62, 0x35, 0x5f, 0x4a, 0x8d, 0x90, 0x32, 0x88, 0xd3, 0x95, 0x18, 0x34, 0x18, 0x9d,
	0x04, 0xa2, 0xd0, 0xbc, 0x9a, 0xc4, 0x06, 0x0e, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x61, 0x2b, 0x65, 0x43, 0x19, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComputeAverageServiceClient is the client API for ComputeAverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComputeAverageServiceClient interface {
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (ComputeAverageService_ComputeAvgClient, error)
}

type computeAverageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeAverageServiceClient(cc grpc.ClientConnInterface) ComputeAverageServiceClient {
	return &computeAverageServiceClient{cc}
}

func (c *computeAverageServiceClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (ComputeAverageService_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ComputeAverageService_serviceDesc.Streams[0], "/computeAverage.ComputeAverageService/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &computeAverageServiceComputeAvgClient{stream}
	return x, nil
}

type ComputeAverageService_ComputeAvgClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type computeAverageServiceComputeAvgClient struct {
	grpc.ClientStream
}

func (x *computeAverageServiceComputeAvgClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *computeAverageServiceComputeAvgClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComputeAverageServiceServer is the server API for ComputeAverageService service.
type ComputeAverageServiceServer interface {
	ComputeAvg(ComputeAverageService_ComputeAvgServer) error
}

// UnimplementedComputeAverageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedComputeAverageServiceServer struct {
}

func (*UnimplementedComputeAverageServiceServer) ComputeAvg(srv ComputeAverageService_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}

func RegisterComputeAverageServiceServer(s *grpc.Server, srv ComputeAverageServiceServer) {
	s.RegisterService(&_ComputeAverageService_serviceDesc, srv)
}

func _ComputeAverageService_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComputeAverageServiceServer).ComputeAvg(&computeAverageServiceComputeAvgServer{stream})
}

type ComputeAverageService_ComputeAvgServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type computeAverageServiceComputeAvgServer struct {
	grpc.ServerStream
}

func (x *computeAverageServiceComputeAvgServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *computeAverageServiceComputeAvgServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ComputeAverageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "computeAverage.ComputeAverageService",
	HandlerType: (*ComputeAverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeAvg",
			Handler:       _ComputeAverageService_ComputeAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "computeAverage.proto",
}
