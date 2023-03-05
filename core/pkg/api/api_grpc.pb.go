// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/api/api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	EntityExists(ctx context.Context, in *EntityExistsRequest, opts ...grpc.CallOption) (*EntityExistsResponse, error)
	CreateEntity(ctx context.Context, in *GenericEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error)
	UpdateEntity(ctx context.Context, in *GenericEntityRequest, opts ...grpc.CallOption) (*UpdateEntityResponse, error)
	AppendEntityHistory(ctx context.Context, in *AppendEntityHistoryRequest, opts ...grpc.CallOption) (*AppendEntityHistoryResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) EntityExists(ctx context.Context, in *EntityExistsRequest, opts ...grpc.CallOption) (*EntityExistsResponse, error) {
	out := new(EntityExistsResponse)
	err := c.cc.Invoke(ctx, "/api/EntityExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreateEntity(ctx context.Context, in *GenericEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error) {
	out := new(CreateEntityResponse)
	err := c.cc.Invoke(ctx, "/api/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateEntity(ctx context.Context, in *GenericEntityRequest, opts ...grpc.CallOption) (*UpdateEntityResponse, error) {
	out := new(UpdateEntityResponse)
	err := c.cc.Invoke(ctx, "/api/UpdateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AppendEntityHistory(ctx context.Context, in *AppendEntityHistoryRequest, opts ...grpc.CallOption) (*AppendEntityHistoryResponse, error) {
	out := new(AppendEntityHistoryResponse)
	err := c.cc.Invoke(ctx, "/api/AppendEntityHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	EntityExists(context.Context, *EntityExistsRequest) (*EntityExistsResponse, error)
	CreateEntity(context.Context, *GenericEntityRequest) (*CreateEntityResponse, error)
	UpdateEntity(context.Context, *GenericEntityRequest) (*UpdateEntityResponse, error)
	AppendEntityHistory(context.Context, *AppendEntityHistoryRequest) (*AppendEntityHistoryResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) EntityExists(context.Context, *EntityExistsRequest) (*EntityExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityExists not implemented")
}
func (UnimplementedApiServer) CreateEntity(context.Context, *GenericEntityRequest) (*CreateEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (UnimplementedApiServer) UpdateEntity(context.Context, *GenericEntityRequest) (*UpdateEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntity not implemented")
}
func (UnimplementedApiServer) AppendEntityHistory(context.Context, *AppendEntityHistoryRequest) (*AppendEntityHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntityHistory not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_EntityExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).EntityExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/EntityExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).EntityExists(ctx, req.(*EntityExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateEntity(ctx, req.(*GenericEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/UpdateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateEntity(ctx, req.(*GenericEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AppendEntityHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntityHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AppendEntityHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/AppendEntityHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AppendEntityHistory(ctx, req.(*AppendEntityHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EntityExists",
			Handler:    _Api_EntityExists_Handler,
		},
		{
			MethodName: "CreateEntity",
			Handler:    _Api_CreateEntity_Handler,
		},
		{
			MethodName: "UpdateEntity",
			Handler:    _Api_UpdateEntity_Handler,
		},
		{
			MethodName: "AppendEntityHistory",
			Handler:    _Api_AppendEntityHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/api.proto",
}
