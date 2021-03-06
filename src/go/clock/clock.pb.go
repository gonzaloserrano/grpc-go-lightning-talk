// Code generated by protoc-gen-go.
// source: clock.proto
// DO NOT EDIT!

/*
Package clock is a generated protocol buffer package.

It is generated from these files:
	clock.proto

It has these top-level messages:
	ClockRequest
	ClockResponse
*/
package clock

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

type ClockRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClockRequest) Reset()                    { *m = ClockRequest{} }
func (m *ClockRequest) String() string            { return proto.CompactTextString(m) }
func (*ClockRequest) ProtoMessage()               {}
func (*ClockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ClockResponse) Reset()                    { *m = ClockResponse{} }
func (m *ClockResponse) String() string            { return proto.CompactTextString(m) }
func (*ClockResponse) ProtoMessage()               {}
func (*ClockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ClockRequest)(nil), "clock.ClockRequest")
	proto.RegisterType((*ClockResponse)(nil), "clock.ClockResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Clock service

type ClockClient interface {
	WhatTimeIsIt(ctx context.Context, in *ClockRequest, opts ...grpc.CallOption) (Clock_WhatTimeIsItClient, error)
}

type clockClient struct {
	cc *grpc.ClientConn
}

func NewClockClient(cc *grpc.ClientConn) ClockClient {
	return &clockClient{cc}
}

func (c *clockClient) WhatTimeIsIt(ctx context.Context, in *ClockRequest, opts ...grpc.CallOption) (Clock_WhatTimeIsItClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Clock_serviceDesc.Streams[0], c.cc, "/clock.Clock/WhatTimeIsIt", opts...)
	if err != nil {
		return nil, err
	}
	x := &clockWhatTimeIsItClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Clock_WhatTimeIsItClient interface {
	Recv() (*ClockResponse, error)
	grpc.ClientStream
}

type clockWhatTimeIsItClient struct {
	grpc.ClientStream
}

func (x *clockWhatTimeIsItClient) Recv() (*ClockResponse, error) {
	m := new(ClockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Clock service

type ClockServer interface {
	WhatTimeIsIt(*ClockRequest, Clock_WhatTimeIsItServer) error
}

func RegisterClockServer(s *grpc.Server, srv ClockServer) {
	s.RegisterService(&_Clock_serviceDesc, srv)
}

func _Clock_WhatTimeIsIt_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClockServer).WhatTimeIsIt(m, &clockWhatTimeIsItServer{stream})
}

type Clock_WhatTimeIsItServer interface {
	Send(*ClockResponse) error
	grpc.ServerStream
}

type clockWhatTimeIsItServer struct {
	grpc.ServerStream
}

func (x *clockWhatTimeIsItServer) Send(m *ClockResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Clock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clock.Clock",
	HandlerType: (*ClockServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WhatTimeIsIt",
			Handler:       _Clock_WhatTimeIsIt_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("clock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xc9, 0x4f,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x94, 0xb8, 0x78, 0x9c,
	0x41, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x93, 0x8b, 0x17, 0xaa, 0xa6,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d,
	0xa6, 0x0e, 0xc6, 0x35, 0x72, 0xe1, 0x62, 0x05, 0x2b, 0x15, 0xb2, 0xe6, 0xe2, 0x09, 0xcf, 0x48,
	0x2c, 0x09, 0xc9, 0xcc, 0x4d, 0xf5, 0x2c, 0xf6, 0x2c, 0x11, 0x12, 0xd6, 0x83, 0x58, 0x8e, 0x6c,
	0x99, 0x94, 0x08, 0xaa, 0x20, 0xc4, 0x74, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x13, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc2, 0x4b, 0x65, 0x0b, 0xb1, 0x00, 0x00, 0x00,
}
