// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1alpha1/cel_service.proto

package expr // import "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CelServiceClient is the client API for CelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CelServiceClient interface {
	// Transforms CEL source text into a parsed representation.
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error)
}

type celServiceClient struct {
	cc *grpc.ClientConn
}

func NewCelServiceClient(cc *grpc.ClientConn) CelServiceClient {
	return &celServiceClient{cc}
}

func (c *celServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error) {
	out := new(EvalResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CelServiceServer is the server API for CelService service.
type CelServiceServer interface {
	// Transforms CEL source text into a parsed representation.
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(context.Context, *EvalRequest) (*EvalResponse, error)
}

func RegisterCelServiceServer(s *grpc.Server, srv CelServiceServer) {
	s.RegisterService(&_CelService_serviceDesc, srv)
}

func _CelService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Eval(ctx, req.(*EvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.api.expr.v1alpha1.CelService",
	HandlerType: (*CelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _CelService_Parse_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CelService_Check_Handler,
		},
		{
			MethodName: "Eval",
			Handler:    _CelService_Eval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/api/expr/v1alpha1/cel_service.proto",
}

func init() {
	proto.RegisterFile("google/api/expr/v1alpha1/cel_service.proto", fileDescriptor_cel_service_082dfec354944354)
}

var fileDescriptor_cel_service_082dfec354944354 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc0, 0x71, 0x2b, 0xea, 0x90, 0x45, 0xc8, 0x24, 0x87, 0x93, 0xe0, 0x09, 0x0e, 0x09, 0x77,
	0x8e, 0x3a, 0xdd, 0xe1, 0x5e, 0x74, 0x10, 0x6e, 0x91, 0x67, 0x78, 0xe6, 0x82, 0x69, 0x5e, 0x4c,
	0x6a, 0xf1, 0xcb, 0xf8, 0x3d, 0x1d, 0x25, 0x69, 0xab, 0x88, 0xc4, 0xde, 0xd8, 0xbe, 0x5f, 0xfe,
	0x81, 0x17, 0x76, 0xa9, 0x89, 0xb4, 0x45, 0x09, 0xde, 0x48, 0x7c, 0xf7, 0x41, 0x76, 0x0b, 0xb0,
	0x7e, 0x0b, 0x0b, 0xa9, 0xd0, 0x3e, 0x46, 0x0c, 0x9d, 0x51, 0x28, 0x7c, 0xa0, 0x96, 0xf8, 0x49,
	0x6f, 0x05, 0x78, 0x23, 0x92, 0x15, 0xa3, 0x9d, 0x2d, 0xcb, 0x15, 0x72, 0xcf, 0x14, 0x1a, 0x70,
	0x0a, 0x7f, 0xd7, 0x96, 0x1f, 0xfb, 0x8c, 0xad, 0xd1, 0xde, 0xf7, 0x3f, 0xf9, 0x86, 0x1d, 0xd6,
	0x10, 0x22, 0xf2, 0xb9, 0x28, 0x5d, 0x23, 0x32, 0xb8, 0xc3, 0xd7, 0x37, 0x8c, 0xed, 0xec, 0x62,
	0xd2, 0x45, 0x4f, 0x2e, 0xe2, 0xd9, 0x5e, 0x6a, 0xaf, 0xb7, 0xa8, 0x5e, 0xfe, 0x6b, 0x67, 0xb0,
	0x43, 0x7b, 0x70, 0xdf, 0xed, 0x07, 0x76, 0x70, 0xdb, 0x81, 0xe5, 0xe7, 0xe5, 0x23, 0x69, 0x3e,
	0x96, 0xe7, 0x53, 0x6c, 0x0c, 0xaf, 0x02, 0x3b, 0x55, 0xd4, 0x14, 0xf9, 0xea, 0xf8, 0x67, 0x79,
	0x75, 0x5a, 0x68, 0x5d, 0x6d, 0x6e, 0x06, 0xac, 0xc9, 0x82, 0xd3, 0x82, 0x82, 0x96, 0x1a, 0x5d,
	0x5e, 0xb7, 0xec, 0x47, 0xe0, 0x4d, 0xfc, 0xfb, 0x4a, 0xd7, 0xe9, 0xeb, 0xb3, 0xaa, 0x9e, 0x8e,
	0xb2, 0xbd, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x97, 0x50, 0xb8, 0x16, 0x02, 0x00, 0x00,
}
