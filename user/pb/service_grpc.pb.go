// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service.proto

package pb

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

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceServiceClient interface {
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	DeleteServiceById(ctx context.Context, in *DeleteServiceByIdRequest, opts ...grpc.CallOption) (*DeleteServiceByIdResponse, error)
	FindServices(ctx context.Context, in *FindServicesRequest, opts ...grpc.CallOption) (*FindServicesResponse, error)
}

type serviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/pb.ServiceService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) DeleteServiceById(ctx context.Context, in *DeleteServiceByIdRequest, opts ...grpc.CallOption) (*DeleteServiceByIdResponse, error) {
	out := new(DeleteServiceByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ServiceService/DeleteServiceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) FindServices(ctx context.Context, in *FindServicesRequest, opts ...grpc.CallOption) (*FindServicesResponse, error) {
	out := new(FindServicesResponse)
	err := c.cc.Invoke(ctx, "/pb.ServiceService/FindServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
// All implementations must embed UnimplementedServiceServiceServer
// for forward compatibility
type ServiceServiceServer interface {
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	DeleteServiceById(context.Context, *DeleteServiceByIdRequest) (*DeleteServiceByIdResponse, error)
	FindServices(context.Context, *FindServicesRequest) (*FindServicesResponse, error)
	mustEmbedUnimplementedServiceServiceServer()
}

// UnimplementedServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (UnimplementedServiceServiceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceServiceServer) DeleteServiceById(context.Context, *DeleteServiceByIdRequest) (*DeleteServiceByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceById not implemented")
}
func (UnimplementedServiceServiceServer) FindServices(context.Context, *FindServicesRequest) (*FindServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindServices not implemented")
}
func (UnimplementedServiceServiceServer) mustEmbedUnimplementedServiceServiceServer() {}

// UnsafeServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServiceServer will
// result in compilation errors.
type UnsafeServiceServiceServer interface {
	mustEmbedUnimplementedServiceServiceServer()
}

func RegisterServiceServiceServer(s grpc.ServiceRegistrar, srv ServiceServiceServer) {
	s.RegisterService(&ServiceService_ServiceDesc, srv)
}

func _ServiceService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServiceService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_DeleteServiceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).DeleteServiceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServiceService/DeleteServiceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).DeleteServiceById(ctx, req.(*DeleteServiceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_FindServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).FindServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServiceService/FindServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).FindServices(ctx, req.(*FindServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceService_ServiceDesc is the grpc.ServiceDesc for ServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _ServiceService_CreateService_Handler,
		},
		{
			MethodName: "DeleteServiceById",
			Handler:    _ServiceService_DeleteServiceById_Handler,
		},
		{
			MethodName: "FindServices",
			Handler:    _ServiceService_FindServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
