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

// WhiteboardServiceClient is the client API for WhiteboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhiteboardServiceClient interface {
	CreateWhiteboard(ctx context.Context, in *CreateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	GetWhiteboard(ctx context.Context, in *GetWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	UpdateWhiteboard(ctx context.Context, in *UpdateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error)
	DeleteWhiteboard(ctx context.Context, in *DeleteWhiteboardRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type whiteboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWhiteboardServiceClient(cc grpc.ClientConnInterface) WhiteboardServiceClient {
	return &whiteboardServiceClient{cc}
}

func (c *whiteboardServiceClient) CreateWhiteboard(ctx context.Context, in *CreateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardService/CreateWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardServiceClient) GetWhiteboard(ctx context.Context, in *GetWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardService/GetWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardServiceClient) UpdateWhiteboard(ctx context.Context, in *UpdateWhiteboardRequest, opts ...grpc.CallOption) (*WhiteboardResponse, error) {
	out := new(WhiteboardResponse)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardService/UpdateWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whiteboardServiceClient) DeleteWhiteboard(ctx context.Context, in *DeleteWhiteboardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/github.com.s1riys.go_whiteboard.WhiteboardService/DeleteWhiteboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhiteboardServiceServer is the server API for WhiteboardService service.
// All implementations must embed UnimplementedWhiteboardServiceServer
// for forward compatibility
type WhiteboardServiceServer interface {
	CreateWhiteboard(context.Context, *CreateWhiteboardRequest) (*WhiteboardResponse, error)
	GetWhiteboard(context.Context, *GetWhiteboardRequest) (*WhiteboardResponse, error)
	UpdateWhiteboard(context.Context, *UpdateWhiteboardRequest) (*WhiteboardResponse, error)
	DeleteWhiteboard(context.Context, *DeleteWhiteboardRequest) (*empty.Empty, error)
	mustEmbedUnimplementedWhiteboardServiceServer()
}

// UnimplementedWhiteboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWhiteboardServiceServer struct {
}

func (UnimplementedWhiteboardServiceServer) CreateWhiteboard(context.Context, *CreateWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWhiteboard not implemented")
}
func (UnimplementedWhiteboardServiceServer) GetWhiteboard(context.Context, *GetWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhiteboard not implemented")
}
func (UnimplementedWhiteboardServiceServer) UpdateWhiteboard(context.Context, *UpdateWhiteboardRequest) (*WhiteboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWhiteboard not implemented")
}
func (UnimplementedWhiteboardServiceServer) DeleteWhiteboard(context.Context, *DeleteWhiteboardRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWhiteboard not implemented")
}
func (UnimplementedWhiteboardServiceServer) mustEmbedUnimplementedWhiteboardServiceServer() {}

// UnsafeWhiteboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhiteboardServiceServer will
// result in compilation errors.
type UnsafeWhiteboardServiceServer interface {
	mustEmbedUnimplementedWhiteboardServiceServer()
}

func RegisterWhiteboardServiceServer(s grpc.ServiceRegistrar, srv WhiteboardServiceServer) {
	s.RegisterService(&WhiteboardService_ServiceDesc, srv)
}

func _WhiteboardService_CreateWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardServiceServer).CreateWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardService/CreateWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardServiceServer).CreateWhiteboard(ctx, req.(*CreateWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardService_GetWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardServiceServer).GetWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardService/GetWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardServiceServer).GetWhiteboard(ctx, req.(*GetWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardService_UpdateWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardServiceServer).UpdateWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardService/UpdateWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardServiceServer).UpdateWhiteboard(ctx, req.(*UpdateWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhiteboardService_DeleteWhiteboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWhiteboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhiteboardServiceServer).DeleteWhiteboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.s1riys.go_whiteboard.WhiteboardService/DeleteWhiteboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhiteboardServiceServer).DeleteWhiteboard(ctx, req.(*DeleteWhiteboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WhiteboardService_ServiceDesc is the grpc.ServiceDesc for WhiteboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhiteboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.s1riys.go_whiteboard.WhiteboardService",
	HandlerType: (*WhiteboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWhiteboard",
			Handler:    _WhiteboardService_CreateWhiteboard_Handler,
		},
		{
			MethodName: "GetWhiteboard",
			Handler:    _WhiteboardService_GetWhiteboard_Handler,
		},
		{
			MethodName: "UpdateWhiteboard",
			Handler:    _WhiteboardService_UpdateWhiteboard_Handler,
		},
		{
			MethodName: "DeleteWhiteboard",
			Handler:    _WhiteboardService_DeleteWhiteboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whiteboard/whiteboard.proto",
}
