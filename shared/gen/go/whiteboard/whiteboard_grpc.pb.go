// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: whiteboard/whiteboard.proto

package whiteboardv1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WhiteboardV1Client is the client API for WhiteboardV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhiteboardV1Client interface {
	CreateWhiteboard(ctx context.Context, in *CreateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	GetWhiteboard(ctx context.Context, in *GetWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	UpdateWhiteboard(ctx context.Context, in *UpdateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	DeleteWhiteboard(ctx context.Context, in *DeleteWhiteboardRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type whiteboardV1Client struct {
	cc grpc.ClientConnInterface
}

func NewWhiteboardV1Client(cc grpc.ClientConnInterface) WhiteboardV1Client {
	return &whiteboardV1Client{cc}
}

func (c *whiteboardV1Client) CreateWhiteboard(ctx context.Context, in *CreateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardV1/CreateWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardV1Client) GetWhiteboard(ctx context.Context, in *GetWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardV1/GetWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardV1Client) UpdateWhiteboard(ctx context.Context, in *UpdateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardV1/UpdateWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardV1Client) DeleteWhiteboard(ctx context.Context, in *DeleteWhiteboardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardV1/DeleteWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhiteboardV1Server is the server API for WhiteboardV1 service.
// All implementations must embed UnimplementedWhiteboardV1Server
// for forward compatibility
type WhiteboardV1Server interface {
	CreateWhiteboard(context.Context, *CreateWhiteboardRequest) (*WhiteboardResponse, error)
	GetWhiteboard(context.Context, *GetWhiteboardRequest) (*WhiteboardResponse, error)
	UpdateWhiteboard(context.Context, *UpdateWhiteboardRequest) (*WhiteboardResponse, error)
	DeleteWhiteboard(context.Context, *DeleteWhiteboardRequest) (*empty.Empty, error)
	mustEmbedUnimplementedWhiteboardV1Server()
}

// UnimplementedWhiteboardV1Server must be embedded to have forward compatible implementations.
type UnimplementedWhiteboardV1Server struct {
}

func (UnimplementedWhiteboardV1Server) CreateWhiteboard(context.Context, *CreateWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWhiteboard not implemented")
}
func (UnimplementedWhiteboardV1Server) GetWhiteboard(context.Context, *GetWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhiteboard not implemented")
}
func (UnimplementedWhiteboardV1Server) UpdateWhiteboard(context.Context, *UpdateWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWhiteboard not implemented")
}
func (UnimplementedWhiteboardV1Server) DeleteWhiteboard(context.Context, *DeleteWhiteboardRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWhiteboard not implemented")
}
func (UnimplementedWhiteboardV1Server) mustEmbedUnimplementedWhiteboardV1Server() {}

// UnsafeWhiteboardV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhiteboardV1Server will
// result in compilation errors.
type UnsafeWhiteboardV1Server interface {
	mustEmbedUnimplementedWhiteboardV1Server()
}

func RegisterWhiteboardV1Server(s grpc.ServiceRegistrar, srv WhiteboardV1Server) {
	s.RegisterService(&WhiteboardV1_ServiceDesc, srv)
}

func _WhiteboardV1_CreateWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardV1Server).CreateWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardV1/CreateWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardV1Server).CreateWhiteboard(ctx, req.(*CreateWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardV1_GetWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardV1Server).GetWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardV1/GetWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardV1Server).GetWhiteboard(ctx, req.(*GetWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardV1_UpdateWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardV1Server).UpdateWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardV1/UpdateWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardV1Server).UpdateWhiteboard(ctx, req.(*UpdateWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardV1_DeleteWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardV1Server).DeleteWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardV1/DeleteWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardV1Server).DeleteWhiteboard(ctx, req.(*DeleteWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WhiteboardV1_ServiceDesc is the grpc.ServiceDesc for WhiteboardV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhiteboardV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.s1riys.go_whiteboard.WhiteboardV1",
	HandlerType: (*WhiteboardV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWhiteboard",
			Handler:    _WhiteboardV1_CreateWhiteboard_Handler,
		},
		{
			MethodName: "GetWhiteboard",
			Handler:    _WhiteboardV1_GetWhiteboard_Handler,
		},
		{
			MethodName: "UpdateWhiteboard",
			Handler:    _WhiteboardV1_UpdateWhiteboard_Handler,
		},
		{
			MethodName: "DeleteWhiteboard",
			Handler:    _WhiteboardV1_DeleteWhiteboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whiteboard/whiteboard.proto",
}
