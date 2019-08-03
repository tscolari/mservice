// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/pb/taxsvc.proto

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

// Add tax request containing the base value
type AddRequest struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d2c9318a915a32, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// The add response containing the value with added tax or an error
type AddReply struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d2c9318a915a32, []int{1}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

func (m *AddReply) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *AddReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// Sub tax request containing the value with the tax
type SubRequest struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubRequest) Reset()         { *m = SubRequest{} }
func (m *SubRequest) String() string { return proto.CompactTextString(m) }
func (*SubRequest) ProtoMessage()    {}
func (*SubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d2c9318a915a32, []int{2}
}

func (m *SubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubRequest.Unmarshal(m, b)
}
func (m *SubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubRequest.Marshal(b, m, deterministic)
}
func (m *SubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubRequest.Merge(m, src)
}
func (m *SubRequest) XXX_Size() int {
	return xxx_messageInfo_SubRequest.Size(m)
}
func (m *SubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubRequest proto.InternalMessageInfo

func (m *SubRequest) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// The sub response containing the value after tax value is subtracted from it
type SubReply struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubReply) Reset()         { *m = SubReply{} }
func (m *SubReply) String() string { return proto.CompactTextString(m) }
func (*SubReply) ProtoMessage()    {}
func (*SubReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d2c9318a915a32, []int{3}
}

func (m *SubReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubReply.Unmarshal(m, b)
}
func (m *SubReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubReply.Marshal(b, m, deterministic)
}
func (m *SubReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubReply.Merge(m, src)
}
func (m *SubReply) XXX_Size() int {
	return xxx_messageInfo_SubReply.Size(m)
}
func (m *SubReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubReply proto.InternalMessageInfo

func (m *SubReply) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SubReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*AddReply)(nil), "pb.AddReply")
	proto.RegisterType((*SubRequest)(nil), "pb.SubRequest")
	proto.RegisterType((*SubReply)(nil), "pb.SubReply")
}

func init() { proto.RegisterFile("pkg/pb/taxsvc.proto", fileDescriptor_36d2c9318a915a32) }

var fileDescriptor_36d2c9318a915a32 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0x4e, 0xd7,
	0x2f, 0x48, 0xd2, 0x2f, 0x49, 0xac, 0x28, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x52, 0x52, 0xe2, 0xe2, 0x72, 0x4c, 0x49, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0a, 0x82, 0x70, 0x94, 0x8c, 0xb8, 0x38, 0xc0, 0x6a, 0x0a, 0x72, 0x2a, 0xb1, 0xab, 0x10, 0x12,
	0xe0, 0x62, 0x4e, 0x2d, 0x2a, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x41, 0xe6,
	0x06, 0x97, 0x26, 0x11, 0x34, 0x17, 0xac, 0x86, 0x04, 0x73, 0x8d, 0x82, 0xb9, 0x98, 0x43, 0x12,
	0x2b, 0x84, 0x54, 0xb9, 0x98, 0x1d, 0x53, 0x52, 0x84, 0xf8, 0xf4, 0x0a, 0x92, 0xf4, 0x10, 0xee,
	0x97, 0xe2, 0x81, 0xf3, 0x0b, 0x72, 0x2a, 0x95, 0x18, 0x40, 0xca, 0x82, 0x4b, 0x93, 0x20, 0xca,
	0x10, 0xce, 0x81, 0x28, 0x83, 0x59, 0xad, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x0f, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0x3b, 0x98, 0x29, 0x26, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaxClient is the client API for Tax service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaxClient interface {
	// Adds the tax value to the base value
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	// Subtracts the tax value from the base value
	Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubReply, error)
}

type taxClient struct {
	cc *grpc.ClientConn
}

func NewTaxClient(cc *grpc.ClientConn) TaxClient {
	return &taxClient{cc}
}

func (c *taxClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/pb.Tax/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxClient) Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubReply, error) {
	out := new(SubReply)
	err := c.cc.Invoke(ctx, "/pb.Tax/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxServer is the server API for Tax service.
type TaxServer interface {
	// Adds the tax value to the base value
	Add(context.Context, *AddRequest) (*AddReply, error)
	// Subtracts the tax value from the base value
	Sub(context.Context, *SubRequest) (*SubReply, error)
}

// UnimplementedTaxServer can be embedded to have forward compatible implementations.
type UnimplementedTaxServer struct {
}

func (*UnimplementedTaxServer) Add(ctx context.Context, req *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTaxServer) Sub(ctx context.Context, req *SubRequest) (*SubReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}

func RegisterTaxServer(s *grpc.Server, srv TaxServer) {
	s.RegisterService(&_Tax_serviceDesc, srv)
}

func _Tax_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tax/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tax_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tax/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxServer).Sub(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tax_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Tax",
	HandlerType: (*TaxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Tax_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Tax_Sub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/taxsvc.proto",
}
