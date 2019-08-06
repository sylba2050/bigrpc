// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package bidirection

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

type Data struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Response struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "bidirection.Data")
	proto.RegisterType((*Response)(nil), "bidirection.Response")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xdc, 0x49, 0x99, 0x29, 0x99, 0x45, 0xa9, 0xc9, 0x25, 0x99, 0xf9,
	0x79, 0x4a, 0x52, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25,
	0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x0c, 0x17, 0x47, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x51, 0x6a, 0x31, 0x54, 0x1a, 0xc4,
	0x34, 0xf2, 0xe1, 0xe2, 0x76, 0x42, 0x18, 0x24, 0x64, 0x8b, 0xca, 0x15, 0xd4, 0x43, 0xb2, 0x45,
	0x0f, 0x64, 0x85, 0x94, 0x28, 0x8a, 0x10, 0xcc, 0x64, 0x25, 0x06, 0x0d, 0x46, 0x03, 0xc6, 0x24,
	0x36, 0xb0, 0xdb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x78, 0x10, 0x59, 0xaa, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BidirectionClient is the client API for Bidirection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BidirectionClient interface {
	Bidirection(ctx context.Context, opts ...grpc.CallOption) (Bidirection_BidirectionClient, error)
}

type bidirectionClient struct {
	cc *grpc.ClientConn
}

func NewBidirectionClient(cc *grpc.ClientConn) BidirectionClient {
	return &bidirectionClient{cc}
}

func (c *bidirectionClient) Bidirection(ctx context.Context, opts ...grpc.CallOption) (Bidirection_BidirectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Bidirection_serviceDesc.Streams[0], "/bidirection.Bidirection/Bidirection", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidirectionBidirectionClient{stream}
	return x, nil
}

type Bidirection_BidirectionClient interface {
	Send(*Data) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type bidirectionBidirectionClient struct {
	grpc.ClientStream
}

func (x *bidirectionBidirectionClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidirectionBidirectionClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionServer is the server API for Bidirection service.
type BidirectionServer interface {
	Bidirection(Bidirection_BidirectionServer) error
}

// UnimplementedBidirectionServer can be embedded to have forward compatible implementations.
type UnimplementedBidirectionServer struct {
}

func (*UnimplementedBidirectionServer) Bidirection(srv Bidirection_BidirectionServer) error {
	return status.Errorf(codes.Unimplemented, "method Bidirection not implemented")
}

func RegisterBidirectionServer(s *grpc.Server, srv BidirectionServer) {
	s.RegisterService(&_Bidirection_serviceDesc, srv)
}

func _Bidirection_Bidirection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidirectionServer).Bidirection(&bidirectionBidirectionServer{stream})
}

type Bidirection_BidirectionServer interface {
	Send(*Response) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type bidirectionBidirectionServer struct {
	grpc.ServerStream
}

func (x *bidirectionBidirectionServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidirectionBidirectionServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Bidirection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bidirection.Bidirection",
	HandlerType: (*BidirectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Bidirection",
			Handler:       _Bidirection_Bidirection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto.proto",
}
