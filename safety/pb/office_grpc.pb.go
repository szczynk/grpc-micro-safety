// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: office.proto

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

// OfficeServiceClient is the client API for OfficeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfficeServiceClient interface {
	CreateOffice(ctx context.Context, in *CreateOfficeRequest, opts ...grpc.CallOption) (*CreateOfficeResponse, error)
	UpdateOfficeById(ctx context.Context, in *UpdateOfficeByIdRequest, opts ...grpc.CallOption) (*UpdateOfficeByIdResponse, error)
	DeleteOfficeById(ctx context.Context, in *DeleteOfficeByIdRequest, opts ...grpc.CallOption) (*DeleteOfficeByIdResponse, error)
	FindOffices(ctx context.Context, in *FindOfficesRequest, opts ...grpc.CallOption) (*FindOfficesResponse, error)
	FindOfficeById(ctx context.Context, in *FindOfficeByIdRequest, opts ...grpc.CallOption) (*FindOfficeByIdResponse, error)
}

type officeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOfficeServiceClient(cc grpc.ClientConnInterface) OfficeServiceClient {
	return &officeServiceClient{cc}
}

func (c *officeServiceClient) CreateOffice(ctx context.Context, in *CreateOfficeRequest, opts ...grpc.CallOption) (*CreateOfficeResponse, error) {
	out := new(CreateOfficeResponse)
	err := c.cc.Invoke(ctx, "/pb.OfficeService/CreateOffice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeServiceClient) UpdateOfficeById(ctx context.Context, in *UpdateOfficeByIdRequest, opts ...grpc.CallOption) (*UpdateOfficeByIdResponse, error) {
	out := new(UpdateOfficeByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.OfficeService/UpdateOfficeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeServiceClient) DeleteOfficeById(ctx context.Context, in *DeleteOfficeByIdRequest, opts ...grpc.CallOption) (*DeleteOfficeByIdResponse, error) {
	out := new(DeleteOfficeByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.OfficeService/DeleteOfficeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeServiceClient) FindOffices(ctx context.Context, in *FindOfficesRequest, opts ...grpc.CallOption) (*FindOfficesResponse, error) {
	out := new(FindOfficesResponse)
	err := c.cc.Invoke(ctx, "/pb.OfficeService/FindOffices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeServiceClient) FindOfficeById(ctx context.Context, in *FindOfficeByIdRequest, opts ...grpc.CallOption) (*FindOfficeByIdResponse, error) {
	out := new(FindOfficeByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.OfficeService/FindOfficeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfficeServiceServer is the server API for OfficeService service.
// All implementations must embed UnimplementedOfficeServiceServer
// for forward compatibility
type OfficeServiceServer interface {
	CreateOffice(context.Context, *CreateOfficeRequest) (*CreateOfficeResponse, error)
	UpdateOfficeById(context.Context, *UpdateOfficeByIdRequest) (*UpdateOfficeByIdResponse, error)
	DeleteOfficeById(context.Context, *DeleteOfficeByIdRequest) (*DeleteOfficeByIdResponse, error)
	FindOffices(context.Context, *FindOfficesRequest) (*FindOfficesResponse, error)
	FindOfficeById(context.Context, *FindOfficeByIdRequest) (*FindOfficeByIdResponse, error)
	mustEmbedUnimplementedOfficeServiceServer()
}

// UnimplementedOfficeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOfficeServiceServer struct {
}

func (UnimplementedOfficeServiceServer) CreateOffice(context.Context, *CreateOfficeRequest) (*CreateOfficeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffice not implemented")
}
func (UnimplementedOfficeServiceServer) UpdateOfficeById(context.Context, *UpdateOfficeByIdRequest) (*UpdateOfficeByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOfficeById not implemented")
}
func (UnimplementedOfficeServiceServer) DeleteOfficeById(context.Context, *DeleteOfficeByIdRequest) (*DeleteOfficeByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOfficeById not implemented")
}
func (UnimplementedOfficeServiceServer) FindOffices(context.Context, *FindOfficesRequest) (*FindOfficesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOffices not implemented")
}
func (UnimplementedOfficeServiceServer) FindOfficeById(context.Context, *FindOfficeByIdRequest) (*FindOfficeByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOfficeById not implemented")
}
func (UnimplementedOfficeServiceServer) mustEmbedUnimplementedOfficeServiceServer() {}

// UnsafeOfficeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfficeServiceServer will
// result in compilation errors.
type UnsafeOfficeServiceServer interface {
	mustEmbedUnimplementedOfficeServiceServer()
}

func RegisterOfficeServiceServer(s grpc.ServiceRegistrar, srv OfficeServiceServer) {
	s.RegisterService(&OfficeService_ServiceDesc, srv)
}

func _OfficeService_CreateOffice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfficeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).CreateOffice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OfficeService/CreateOffice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).CreateOffice(ctx, req.(*CreateOfficeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeService_UpdateOfficeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOfficeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).UpdateOfficeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OfficeService/UpdateOfficeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).UpdateOfficeById(ctx, req.(*UpdateOfficeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeService_DeleteOfficeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOfficeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).DeleteOfficeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OfficeService/DeleteOfficeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).DeleteOfficeById(ctx, req.(*DeleteOfficeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeService_FindOffices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOfficesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).FindOffices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OfficeService/FindOffices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).FindOffices(ctx, req.(*FindOfficesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeService_FindOfficeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOfficeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).FindOfficeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OfficeService/FindOfficeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).FindOfficeById(ctx, req.(*FindOfficeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OfficeService_ServiceDesc is the grpc.ServiceDesc for OfficeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OfficeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OfficeService",
	HandlerType: (*OfficeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOffice",
			Handler:    _OfficeService_CreateOffice_Handler,
		},
		{
			MethodName: "UpdateOfficeById",
			Handler:    _OfficeService_UpdateOfficeById_Handler,
		},
		{
			MethodName: "DeleteOfficeById",
			Handler:    _OfficeService_DeleteOfficeById_Handler,
		},
		{
			MethodName: "FindOffices",
			Handler:    _OfficeService_FindOffices_Handler,
		},
		{
			MethodName: "FindOfficeById",
			Handler:    _OfficeService_FindOfficeById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "office.proto",
}
