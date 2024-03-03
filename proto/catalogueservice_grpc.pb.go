// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: catalogueservice.proto

package catalogueservice

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

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	Create(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	Get(ctx context.Context, in *RestaurantIdRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	GetAll(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*RestaurantListResponse, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) Create(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/RestaurantService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) Get(ctx context.Context, in *RestaurantIdRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/RestaurantService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetAll(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*RestaurantListResponse, error) {
	out := new(RestaurantListResponse)
	err := c.cc.Invoke(ctx, "/RestaurantService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility
type RestaurantServiceServer interface {
	Create(context.Context, *RestaurantRequest) (*RestaurantResponse, error)
	Get(context.Context, *RestaurantIdRequest) (*RestaurantResponse, error)
	GetAll(context.Context, *NoParams) (*RestaurantListResponse, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantServiceServer struct {
}

func (UnimplementedRestaurantServiceServer) Create(context.Context, *RestaurantRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRestaurantServiceServer) Get(context.Context, *RestaurantIdRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRestaurantServiceServer) GetAll(context.Context, *NoParams) (*RestaurantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RestaurantService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Create(ctx, req.(*RestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RestaurantService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).Get(ctx, req.(*RestaurantIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RestaurantService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetAll(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RestaurantService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RestaurantService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RestaurantService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalogueservice.proto",
}
