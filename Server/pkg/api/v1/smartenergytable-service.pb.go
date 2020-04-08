// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smartenergytable-service.proto

package v1

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("smartenergytable-service.proto", fileDescriptor_7be9bf9d675643a6) }

var fileDescriptor_7be9bf9d675643a6 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xce, 0x4d, 0x2c,
	0x2a, 0x49, 0xcd, 0x4b, 0x2d, 0x4a, 0xaf, 0x2c, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2d, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe7, 0x62, 0x75, 0xcd,
	0x2d, 0x28, 0xa9, 0x34, 0x32, 0xe7, 0x12, 0x0f, 0x06, 0x29, 0x75, 0x05, 0x2b, 0x0d, 0x01, 0x29,
	0x0d, 0x86, 0xa8, 0x14, 0x92, 0xe1, 0x62, 0x0b, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x15, 0x62, 0xd3,
	0x03, 0x2b, 0x96, 0x82, 0xd2, 0x4a, 0x0c, 0x06, 0x8c, 0x4e, 0x3c, 0x51, 0x5c, 0x05, 0xd9, 0xe9,
	0xfa, 0x89, 0x05, 0x99, 0xfa, 0x65, 0x86, 0x49, 0x6c, 0x60, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0xb7, 0xa8, 0xa5, 0x78, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmartEnergyTableServiceClient is the client API for SmartEnergyTableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmartEnergyTableServiceClient interface {
	Update(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SmartEnergyTableService_UpdateClient, error)
}

type smartEnergyTableServiceClient struct {
	cc *grpc.ClientConn
}

func NewSmartEnergyTableServiceClient(cc *grpc.ClientConn) SmartEnergyTableServiceClient {
	return &smartEnergyTableServiceClient{cc}
}

func (c *smartEnergyTableServiceClient) Update(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SmartEnergyTableService_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmartEnergyTableService_serviceDesc.Streams[0], "/SmartEnergyTableService/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartEnergyTableServiceUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartEnergyTableService_UpdateClient interface {
	Recv() (*Empty, error)
	grpc.ClientStream
}

type smartEnergyTableServiceUpdateClient struct {
	grpc.ClientStream
}

func (x *smartEnergyTableServiceUpdateClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SmartEnergyTableServiceServer is the server API for SmartEnergyTableService service.
type SmartEnergyTableServiceServer interface {
	Update(*Empty, SmartEnergyTableService_UpdateServer) error
}

// UnimplementedSmartEnergyTableServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmartEnergyTableServiceServer struct {
}

func (*UnimplementedSmartEnergyTableServiceServer) Update(req *Empty, srv SmartEnergyTableService_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterSmartEnergyTableServiceServer(s *grpc.Server, srv SmartEnergyTableServiceServer) {
	s.RegisterService(&_SmartEnergyTableService_serviceDesc, srv)
}

func _SmartEnergyTableService_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartEnergyTableServiceServer).Update(m, &smartEnergyTableServiceUpdateServer{stream})
}

type SmartEnergyTableService_UpdateServer interface {
	Send(*Empty) error
	grpc.ServerStream
}

type smartEnergyTableServiceUpdateServer struct {
	grpc.ServerStream
}

func (x *smartEnergyTableServiceUpdateServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

var _SmartEnergyTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SmartEnergyTableService",
	HandlerType: (*SmartEnergyTableServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Update",
			Handler:       _SmartEnergyTableService_Update_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "smartenergytable-service.proto",
}
