// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: certificate.proto

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

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CreateCertificateResponse, error)
	UpdateCertificateById(ctx context.Context, in *UpdateCertificateByIdRequest, opts ...grpc.CallOption) (*UpdateCertificateByIdResponse, error)
	DeleteCertificateById(ctx context.Context, in *DeleteCertificateByIdRequest, opts ...grpc.CallOption) (*DeleteCertificateByIdResponse, error)
	FindCertificates(ctx context.Context, in *FindCertificatesRequest, opts ...grpc.CallOption) (*FindCertificatesResponse, error)
	FindCertificateById(ctx context.Context, in *FindCertificateByIdRequest, opts ...grpc.CallOption) (*FindCertificateByIdResponse, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CreateCertificateResponse, error) {
	out := new(CreateCertificateResponse)
	err := c.cc.Invoke(ctx, "/pb.CertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) UpdateCertificateById(ctx context.Context, in *UpdateCertificateByIdRequest, opts ...grpc.CallOption) (*UpdateCertificateByIdResponse, error) {
	out := new(UpdateCertificateByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.CertificateService/UpdateCertificateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) DeleteCertificateById(ctx context.Context, in *DeleteCertificateByIdRequest, opts ...grpc.CallOption) (*DeleteCertificateByIdResponse, error) {
	out := new(DeleteCertificateByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.CertificateService/DeleteCertificateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) FindCertificates(ctx context.Context, in *FindCertificatesRequest, opts ...grpc.CallOption) (*FindCertificatesResponse, error) {
	out := new(FindCertificatesResponse)
	err := c.cc.Invoke(ctx, "/pb.CertificateService/FindCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) FindCertificateById(ctx context.Context, in *FindCertificateByIdRequest, opts ...grpc.CallOption) (*FindCertificateByIdResponse, error) {
	out := new(FindCertificateByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.CertificateService/FindCertificateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility
type CertificateServiceServer interface {
	CreateCertificate(context.Context, *CreateCertificateRequest) (*CreateCertificateResponse, error)
	UpdateCertificateById(context.Context, *UpdateCertificateByIdRequest) (*UpdateCertificateByIdResponse, error)
	DeleteCertificateById(context.Context, *DeleteCertificateByIdRequest) (*DeleteCertificateByIdResponse, error)
	FindCertificates(context.Context, *FindCertificatesRequest) (*FindCertificatesResponse, error)
	FindCertificateById(context.Context, *FindCertificateByIdRequest) (*FindCertificateByIdResponse, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (UnimplementedCertificateServiceServer) CreateCertificate(context.Context, *CreateCertificateRequest) (*CreateCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) UpdateCertificateById(context.Context, *UpdateCertificateByIdRequest) (*UpdateCertificateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCertificateById not implemented")
}
func (UnimplementedCertificateServiceServer) DeleteCertificateById(context.Context, *DeleteCertificateByIdRequest) (*DeleteCertificateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificateById not implemented")
}
func (UnimplementedCertificateServiceServer) FindCertificates(context.Context, *FindCertificatesRequest) (*FindCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCertificates not implemented")
}
func (UnimplementedCertificateServiceServer) FindCertificateById(context.Context, *FindCertificateByIdRequest) (*FindCertificateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCertificateById not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*CreateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_UpdateCertificateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).UpdateCertificateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertificateService/UpdateCertificateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).UpdateCertificateById(ctx, req.(*UpdateCertificateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_DeleteCertificateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).DeleteCertificateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertificateService/DeleteCertificateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).DeleteCertificateById(ctx, req.(*DeleteCertificateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_FindCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).FindCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertificateService/FindCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).FindCertificates(ctx, req.(*FindCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_FindCertificateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCertificateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).FindCertificateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertificateService/FindCertificateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).FindCertificateById(ctx, req.(*FindCertificateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
		{
			MethodName: "UpdateCertificateById",
			Handler:    _CertificateService_UpdateCertificateById_Handler,
		},
		{
			MethodName: "DeleteCertificateById",
			Handler:    _CertificateService_DeleteCertificateById_Handler,
		},
		{
			MethodName: "FindCertificates",
			Handler:    _CertificateService_FindCertificates_Handler,
		},
		{
			MethodName: "FindCertificateById",
			Handler:    _CertificateService_FindCertificateById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}
