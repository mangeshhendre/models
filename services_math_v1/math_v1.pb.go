// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/math/math_v1.proto

/*
Package services_math_v1 is a generated protocol buffer package.

It is generated from these files:
	services/math/math_v1.proto

It has these top-level messages:
	MathRequest
	MathResponse
*/
package services_math_v1

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

type MathRequest struct {
	Number1 float32 `protobuf:"fixed32,1,opt,name=number1" json:"number1,omitempty"`
	Number2 float32 `protobuf:"fixed32,2,opt,name=number2" json:"number2,omitempty"`
}

func (m *MathRequest) Reset()                    { *m = MathRequest{} }
func (m *MathRequest) String() string            { return proto.CompactTextString(m) }
func (*MathRequest) ProtoMessage()               {}
func (*MathRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MathRequest) GetNumber1() float32 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *MathRequest) GetNumber2() float32 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type MathResponse struct {
	Result float32 `protobuf:"fixed32,1,opt,name=result" json:"result,omitempty"`
}

func (m *MathResponse) Reset()                    { *m = MathResponse{} }
func (m *MathResponse) String() string            { return proto.CompactTextString(m) }
func (*MathResponse) ProtoMessage()               {}
func (*MathResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MathResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*MathRequest)(nil), "services.luggage.v1.MathRequest")
	proto.RegisterType((*MathResponse)(nil), "services.luggage.v1.MathResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Math service

type MathClient interface {
	AddNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	MultiplyNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	DevideNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
}

type mathClient struct {
	cc *grpc.ClientConn
}

func NewMathClient(cc *grpc.ClientConn) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) AddNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := grpc.Invoke(ctx, "/services.luggage.v1.Math/AddNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) MultiplyNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := grpc.Invoke(ctx, "/services.luggage.v1.Math/MultiplyNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) DevideNumber(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := grpc.Invoke(ctx, "/services.luggage.v1.Math/DevideNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Math service

type MathServer interface {
	AddNumber(context.Context, *MathRequest) (*MathResponse, error)
	MultiplyNumber(context.Context, *MathRequest) (*MathResponse, error)
	DevideNumber(context.Context, *MathRequest) (*MathResponse, error)
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_AddNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).AddNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.luggage.v1.Math/AddNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).AddNumber(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_MultiplyNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).MultiplyNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.luggage.v1.Math/MultiplyNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).MultiplyNumber(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_DevideNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).DevideNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.luggage.v1.Math/DevideNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).DevideNumber(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.luggage.v1.Math",
	HandlerType: (*MathServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumber",
			Handler:    _Math_AddNumber_Handler,
		},
		{
			MethodName: "MultiplyNumber",
			Handler:    _Math_MultiplyNumber_Handler,
		},
		{
			MethodName: "DevideNumber",
			Handler:    _Math_DevideNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/math/math_v1.proto",
}

func init() { proto.RegisterFile("services/math/math_v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0xcf, 0x4d, 0x2c, 0xc9, 0x00, 0x13, 0xf1, 0x65, 0x86, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xc2, 0x30, 0x49, 0xbd, 0x9c, 0xd2, 0xf4, 0xf4, 0xc4, 0xf4, 0x54,
	0xbd, 0x32, 0x43, 0x25, 0x47, 0x2e, 0x6e, 0xdf, 0xc4, 0x92, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x43, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xa6, 0x20, 0x18, 0x17, 0x21, 0x63, 0x24, 0xc1, 0x84, 0x2c, 0x63, 0xa4, 0xa4, 0xc6,
	0xc5, 0x03, 0x31, 0xa2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5,
	0xb8, 0x34, 0xa7, 0x04, 0x6a, 0x04, 0x94, 0x67, 0xd4, 0xc2, 0xc4, 0xc5, 0x02, 0x52, 0x28, 0x14,
	0xc0, 0xc5, 0xe9, 0x98, 0x92, 0xe2, 0x07, 0xd6, 0x2e, 0xa4, 0xa0, 0x87, 0xc5, 0x59, 0x7a, 0x48,
	0x6e, 0x92, 0x52, 0xc4, 0xa3, 0x02, 0x6a, 0x65, 0x28, 0x17, 0x9f, 0x6f, 0x69, 0x4e, 0x49, 0x66,
	0x41, 0x4e, 0x25, 0x35, 0x8d, 0x0d, 0xe6, 0xe2, 0x71, 0x49, 0x2d, 0xcb, 0x4c, 0x49, 0xa5, 0xa2,
	0xa1, 0x4e, 0x46, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0xb9, 0x89, 0x79, 0xe9, 0xa9, 0xc5, 0x19, 0x19, 0xa9, 0x79, 0x29, 0x45, 0xa9, 0xfa, 0xb9, 0xf9,
	0x29, 0xa9, 0x39, 0xc5, 0xfa, 0x30, 0x33, 0xe2, 0xa1, 0x11, 0x98, 0xc4, 0x06, 0x8e, 0x41, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xd8, 0xca, 0xd0, 0xe0, 0x01, 0x00, 0x00,
}
