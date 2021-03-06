// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defaultcmd.proto

package cmdpb

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

func init() { proto.RegisterFile("defaultcmd.proto", fileDescriptor_4f24e0a83d0b1c21) }

var fileDescriptor_4f24e0a83d0b1c21 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x4d, 0x4b,
	0x2c, 0xcd, 0x29, 0x49, 0xce, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xce,
	0x4d, 0x29, 0x48, 0x92, 0xe2, 0x85, 0x4a, 0x40, 0x44, 0x8d, 0x7c, 0xb8, 0x78, 0x11, 0x2a, 0x8b,
	0x8b, 0xca, 0x84, 0xac, 0xb9, 0x78, 0x5c, 0x20, 0x02, 0xce, 0xb9, 0x29, 0x2e, 0xf9, 0x42, 0xa2,
	0x7a, 0x60, 0x7d, 0x7a, 0x50, 0xc1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x21, 0x74,
	0xe1, 0xe2, 0x02, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0xd1, 0x1b, 0x5f, 0x7e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DefaultcmdsrvClient is the client API for Defaultcmdsrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DefaultcmdsrvClient interface {
	DefaultCmdDo(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*DefaultResp, error)
}

type defaultcmdsrvClient struct {
	cc *grpc.ClientConn
}

func NewDefaultcmdsrvClient(cc *grpc.ClientConn) DefaultcmdsrvClient {
	return &defaultcmdsrvClient{cc}
}

func (c *defaultcmdsrvClient) DefaultCmdDo(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*DefaultResp, error) {
	out := new(DefaultResp)
	err := c.cc.Invoke(ctx, "/cmdpb.defaultcmdsrv/DefaultCmdDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultcmdsrvServer is the server API for Defaultcmdsrv service.
type DefaultcmdsrvServer interface {
	DefaultCmdDo(context.Context, *DefaultRequest) (*DefaultResp, error)
}

// UnimplementedDefaultcmdsrvServer can be embedded to have forward compatible implementations.
type UnimplementedDefaultcmdsrvServer struct {
}

func (*UnimplementedDefaultcmdsrvServer) DefaultCmdDo(ctx context.Context, req *DefaultRequest) (*DefaultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultCmdDo not implemented")
}

func RegisterDefaultcmdsrvServer(s *grpc.Server, srv DefaultcmdsrvServer) {
	s.RegisterService(&_Defaultcmdsrv_serviceDesc, srv)
}

func _Defaultcmdsrv_DefaultCmdDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultcmdsrvServer).DefaultCmdDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmdpb.defaultcmdsrv/DefaultCmdDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultcmdsrvServer).DefaultCmdDo(ctx, req.(*DefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Defaultcmdsrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cmdpb.defaultcmdsrv",
	HandlerType: (*DefaultcmdsrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DefaultCmdDo",
			Handler:    _Defaultcmdsrv_DefaultCmdDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "defaultcmd.proto",
}
