// Code generated by protoc-gen-go. DO NOT EDIT.
// source: world.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SayWorldRequest struct {
	YourWold             string   `protobuf:"bytes,1,opt,name=yourWold,proto3" json:"yourWold,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayWorldRequest) Reset()         { *m = SayWorldRequest{} }
func (m *SayWorldRequest) String() string { return proto.CompactTextString(m) }
func (*SayWorldRequest) ProtoMessage()    {}
func (*SayWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_world_23e35526575181a8, []int{0}
}
func (m *SayWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayWorldRequest.Unmarshal(m, b)
}
func (m *SayWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayWorldRequest.Marshal(b, m, deterministic)
}
func (dst *SayWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayWorldRequest.Merge(dst, src)
}
func (m *SayWorldRequest) XXX_Size() int {
	return xxx_messageInfo_SayWorldRequest.Size(m)
}
func (m *SayWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayWorldRequest proto.InternalMessageInfo

func (m *SayWorldRequest) GetYourWold() string {
	if m != nil {
		return m.YourWold
	}
	return ""
}

func (m *SayWorldRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SayWorldResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayWorldResponse) Reset()         { *m = SayWorldResponse{} }
func (m *SayWorldResponse) String() string { return proto.CompactTextString(m) }
func (*SayWorldResponse) ProtoMessage()    {}
func (*SayWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_world_23e35526575181a8, []int{1}
}
func (m *SayWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayWorldResponse.Unmarshal(m, b)
}
func (m *SayWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayWorldResponse.Marshal(b, m, deterministic)
}
func (dst *SayWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayWorldResponse.Merge(dst, src)
}
func (m *SayWorldResponse) XXX_Size() int {
	return xxx_messageInfo_SayWorldResponse.Size(m)
}
func (m *SayWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayWorldResponse proto.InternalMessageInfo

func (m *SayWorldResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayWorldRequest)(nil), "proto.SayWorldRequest")
	proto.RegisterType((*SayWorldResponse)(nil), "proto.SayWorldResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorldClient is the client API for World service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorldClient interface {
	SayWorld(ctx context.Context, in *SayWorldRequest, opts ...grpc.CallOption) (*SayWorldResponse, error)
}

type worldClient struct {
	cc *grpc.ClientConn
}

func NewWorldClient(cc *grpc.ClientConn) WorldClient {
	return &worldClient{cc}
}

func (c *worldClient) SayWorld(ctx context.Context, in *SayWorldRequest, opts ...grpc.CallOption) (*SayWorldResponse, error) {
	out := new(SayWorldResponse)
	err := c.cc.Invoke(ctx, "/proto.World/sayWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorldServer is the server API for World service.
type WorldServer interface {
	SayWorld(context.Context, *SayWorldRequest) (*SayWorldResponse, error)
}

func RegisterWorldServer(s *grpc.Server, srv WorldServer) {
	s.RegisterService(&_World_serviceDesc, srv)
}

func _World_SayWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldServer).SayWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.World/SayWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldServer).SayWorld(ctx, req.(*SayWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _World_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.World",
	HandlerType: (*WorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayWorld",
			Handler:    _World_SayWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "world.proto",
}

func init() { proto.RegisterFile("world.proto", fileDescriptor_world_23e35526575181a8) }

var fileDescriptor_world_23e35526575181a8 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xcf, 0x2f, 0xca,
	0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x9e, 0x5c, 0xfc, 0xc1,
	0x89, 0x95, 0xe1, 0x20, 0x89, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e,
	0xca, 0xfc, 0xd2, 0xa2, 0xf0, 0xfc, 0x9c, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38,
	0x1f, 0x24, 0x57, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1, 0x04, 0x91, 0x83, 0xf1,
	0x95, 0x74, 0xb8, 0x04, 0x10, 0x46, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1,
	0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x8d, 0x82, 0x71, 0x8d, 0xdc, 0xb8, 0x58, 0xc1,
	0x4a, 0x85, 0x6c, 0xb9, 0x38, 0x8a, 0xa1, 0xda, 0x84, 0xc4, 0x20, 0x8e, 0xd3, 0x43, 0x73, 0x92,
	0x94, 0x38, 0x86, 0x38, 0xc4, 0x7c, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x8c, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x3f, 0x4b, 0x26, 0xdd, 0x00, 0x00, 0x00,
}