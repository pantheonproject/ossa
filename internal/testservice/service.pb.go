// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package main

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

type Thing struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DataPoint            int64    `protobuf:"varint,3,opt,name=data_point,json=dataPoint,proto3" json:"data_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thing) Reset()         { *m = Thing{} }
func (m *Thing) String() string { return proto.CompactTextString(m) }
func (*Thing) ProtoMessage()    {}
func (*Thing) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Thing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thing.Unmarshal(m, b)
}
func (m *Thing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thing.Marshal(b, m, deterministic)
}
func (m *Thing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thing.Merge(m, src)
}
func (m *Thing) XXX_Size() int {
	return xxx_messageInfo_Thing.Size(m)
}
func (m *Thing) XXX_DiscardUnknown() {
	xxx_messageInfo_Thing.DiscardUnknown(m)
}

var xxx_messageInfo_Thing proto.InternalMessageInfo

func (m *Thing) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Thing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thing) GetDataPoint() int64 {
	if m != nil {
		return m.DataPoint
	}
	return 0
}

type AddThingRequest struct {
	Thing                *Thing   `protobuf:"bytes,1,opt,name=thing,proto3" json:"thing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddThingRequest) Reset()         { *m = AddThingRequest{} }
func (m *AddThingRequest) String() string { return proto.CompactTextString(m) }
func (*AddThingRequest) ProtoMessage()    {}
func (*AddThingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *AddThingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddThingRequest.Unmarshal(m, b)
}
func (m *AddThingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddThingRequest.Marshal(b, m, deterministic)
}
func (m *AddThingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddThingRequest.Merge(m, src)
}
func (m *AddThingRequest) XXX_Size() int {
	return xxx_messageInfo_AddThingRequest.Size(m)
}
func (m *AddThingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddThingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddThingRequest proto.InternalMessageInfo

func (m *AddThingRequest) GetThing() *Thing {
	if m != nil {
		return m.Thing
	}
	return nil
}

type AddThingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddThingResponse) Reset()         { *m = AddThingResponse{} }
func (m *AddThingResponse) String() string { return proto.CompactTextString(m) }
func (*AddThingResponse) ProtoMessage()    {}
func (*AddThingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *AddThingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddThingResponse.Unmarshal(m, b)
}
func (m *AddThingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddThingResponse.Marshal(b, m, deterministic)
}
func (m *AddThingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddThingResponse.Merge(m, src)
}
func (m *AddThingResponse) XXX_Size() int {
	return xxx_messageInfo_AddThingResponse.Size(m)
}
func (m *AddThingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddThingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddThingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Thing)(nil), "main.Thing")
	proto.RegisterType((*AddThingRequest)(nil), "main.AddThingRequest")
	proto.RegisterType((*AddThingResponse)(nil), "main.AddThingResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93,
	0xe2, 0x4e, 0x2c, 0x4d, 0xc9, 0x2c, 0x81, 0x08, 0x29, 0x05, 0x70, 0xb1, 0x86, 0x64, 0x64, 0xe6,
	0xa5, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65,
	0xa6, 0x08, 0x49, 0x70, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x44, 0x9c, 0x58, 0x3e,
	0xcc, 0x62, 0x66, 0x0c, 0x02, 0x8b, 0x08, 0xc9, 0x72, 0x71, 0xa5, 0x24, 0x96, 0x24, 0xc6, 0x17,
	0xe4, 0x67, 0xe6, 0x95, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x82, 0x44, 0x02, 0x40,
	0x02, 0x4a, 0x26, 0x5c, 0xfc, 0x8e, 0x29, 0x29, 0x60, 0x43, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x14, 0xb9, 0x58, 0x4b, 0x40, 0x7c, 0xb0, 0xf1, 0xdc, 0x46, 0xdc, 0x7a, 0x20, 0x77,
	0xe8, 0x41, 0x94, 0x40, 0x64, 0x94, 0x84, 0xb8, 0x04, 0x10, 0xba, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x8d, 0x1c, 0xb9, 0x58, 0x42, 0x40, 0xda, 0x2d, 0xb9, 0x38, 0x60, 0x72, 0x42, 0xa2, 0x10,
	0xbd, 0x68, 0x36, 0x48, 0x89, 0xa1, 0x0b, 0x43, 0x8c, 0x48, 0x62, 0x03, 0xfb, 0xd2, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x92, 0xb1, 0x89, 0xc2, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	AddThing(ctx context.Context, in *AddThingRequest, opts ...grpc.CallOption) (*AddThingResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) AddThing(ctx context.Context, in *AddThingRequest, opts ...grpc.CallOption) (*AddThingResponse, error) {
	out := new(AddThingResponse)
	err := c.cc.Invoke(ctx, "/main.Test/AddThing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	AddThing(context.Context, *AddThingRequest) (*AddThingResponse, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) AddThing(ctx context.Context, req *AddThingRequest) (*AddThingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddThing not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_AddThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddThingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).AddThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Test/AddThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).AddThing(ctx, req.(*AddThingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddThing",
			Handler:    _Test_AddThing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
