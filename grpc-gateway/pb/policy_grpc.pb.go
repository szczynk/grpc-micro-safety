// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: policy.proto

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

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyServiceClient interface {
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error)
	FindPolicies(ctx context.Context, in *FindPoliciesRequest, opts ...grpc.CallOption) (*FindPoliciesResponse, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error) {
	out := new(CreatePolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.PolicyService/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error) {
	out := new(DeletePolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.PolicyService/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) FindPolicies(ctx context.Context, in *FindPoliciesRequest, opts ...grpc.CallOption) (*FindPoliciesResponse, error) {
	out := new(FindPoliciesResponse)
	err := c.cc.Invoke(ctx, "/pb.PolicyService/FindPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
// All implementations must embed UnimplementedPolicyServiceServer
// for forward compatibility
type PolicyServiceServer interface {
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)
	FindPolicies(context.Context, *FindPoliciesRequest) (*FindPoliciesResponse, error)
	mustEmbedUnimplementedPolicyServiceServer()
}

// UnimplementedPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (UnimplementedPolicyServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) FindPolicies(context.Context, *FindPoliciesRequest) (*FindPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPolicies not implemented")
}
func (UnimplementedPolicyServiceServer) mustEmbedUnimplementedPolicyServiceServer() {}

// UnsafePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServiceServer will
// result in compilation errors.
type UnsafePolicyServiceServer interface {
	mustEmbedUnimplementedPolicyServiceServer()
}

func RegisterPolicyServiceServer(s grpc.ServiceRegistrar, srv PolicyServiceServer) {
	s.RegisterService(&PolicyService_ServiceDesc, srv)
}

func _PolicyService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PolicyService/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PolicyService/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_FindPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).FindPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PolicyService/FindPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).FindPolicies(ctx, req.(*FindPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyService_ServiceDesc is the grpc.ServiceDesc for PolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePolicy",
			Handler:    _PolicyService_CreatePolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _PolicyService_DeletePolicy_Handler,
		},
		{
			MethodName: "FindPolicies",
			Handler:    _PolicyService_FindPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy.proto",
}