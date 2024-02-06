// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: objects.proto

package proto

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

// DomainServiceClient is the client API for DomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainServiceClient interface {
	GetDomainsPool(ctx context.Context, in *GetDomainsPoolRequest, opts ...grpc.CallOption) (*GetDomainsPoolResponse, error)
}

type domainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainServiceClient(cc grpc.ClientConnInterface) DomainServiceClient {
	return &domainServiceClient{cc}
}

func (c *domainServiceClient) GetDomainsPool(ctx context.Context, in *GetDomainsPoolRequest, opts ...grpc.CallOption) (*GetDomainsPoolResponse, error) {
	out := new(GetDomainsPoolResponse)
	err := c.cc.Invoke(ctx, "/objects.DomainService/GetDomainsPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServiceServer is the server API for DomainService service.
// All implementations must embed UnimplementedDomainServiceServer
// for forward compatibility
type DomainServiceServer interface {
	GetDomainsPool(context.Context, *GetDomainsPoolRequest) (*GetDomainsPoolResponse, error)
	mustEmbedUnimplementedDomainServiceServer()
}

// UnimplementedDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainServiceServer struct {
}

func (UnimplementedDomainServiceServer) GetDomainsPool(context.Context, *GetDomainsPoolRequest) (*GetDomainsPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainsPool not implemented")
}
func (UnimplementedDomainServiceServer) mustEmbedUnimplementedDomainServiceServer() {}

// UnsafeDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainServiceServer will
// result in compilation errors.
type UnsafeDomainServiceServer interface {
	mustEmbedUnimplementedDomainServiceServer()
}

func RegisterDomainServiceServer(s grpc.ServiceRegistrar, srv DomainServiceServer) {
	s.RegisterService(&DomainService_ServiceDesc, srv)
}

func _DomainService_GetDomainsPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainsPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).GetDomainsPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objects.DomainService/GetDomainsPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).GetDomainsPool(ctx, req.(*GetDomainsPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainService_ServiceDesc is the grpc.ServiceDesc for DomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "objects.DomainService",
	HandlerType: (*DomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomainsPool",
			Handler:    _DomainService_GetDomainsPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objects.proto",
}

// SubdomainServiceClient is the client API for SubdomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubdomainServiceClient interface {
	GetOccupiedSubdomains(ctx context.Context, in *GetOccupiedSubdomainsRequest, opts ...grpc.CallOption) (*GetOccupiedSubdomainsResponse, error)
}

type subdomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubdomainServiceClient(cc grpc.ClientConnInterface) SubdomainServiceClient {
	return &subdomainServiceClient{cc}
}

func (c *subdomainServiceClient) GetOccupiedSubdomains(ctx context.Context, in *GetOccupiedSubdomainsRequest, opts ...grpc.CallOption) (*GetOccupiedSubdomainsResponse, error) {
	out := new(GetOccupiedSubdomainsResponse)
	err := c.cc.Invoke(ctx, "/objects.SubdomainService/GetOccupiedSubdomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubdomainServiceServer is the server API for SubdomainService service.
// All implementations must embed UnimplementedSubdomainServiceServer
// for forward compatibility
type SubdomainServiceServer interface {
	GetOccupiedSubdomains(context.Context, *GetOccupiedSubdomainsRequest) (*GetOccupiedSubdomainsResponse, error)
	mustEmbedUnimplementedSubdomainServiceServer()
}

// UnimplementedSubdomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubdomainServiceServer struct {
}

func (UnimplementedSubdomainServiceServer) GetOccupiedSubdomains(context.Context, *GetOccupiedSubdomainsRequest) (*GetOccupiedSubdomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOccupiedSubdomains not implemented")
}
func (UnimplementedSubdomainServiceServer) mustEmbedUnimplementedSubdomainServiceServer() {}

// UnsafeSubdomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubdomainServiceServer will
// result in compilation errors.
type UnsafeSubdomainServiceServer interface {
	mustEmbedUnimplementedSubdomainServiceServer()
}

func RegisterSubdomainServiceServer(s grpc.ServiceRegistrar, srv SubdomainServiceServer) {
	s.RegisterService(&SubdomainService_ServiceDesc, srv)
}

func _SubdomainService_GetOccupiedSubdomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOccupiedSubdomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubdomainServiceServer).GetOccupiedSubdomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objects.SubdomainService/GetOccupiedSubdomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubdomainServiceServer).GetOccupiedSubdomains(ctx, req.(*GetOccupiedSubdomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubdomainService_ServiceDesc is the grpc.ServiceDesc for SubdomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubdomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "objects.SubdomainService",
	HandlerType: (*SubdomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOccupiedSubdomains",
			Handler:    _SubdomainService_GetOccupiedSubdomains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objects.proto",
}

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	AddRouteToProject(ctx context.Context, in *AddRouteToProjectRequest, opts ...grpc.CallOption) (*AddRouteToProjectResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, "/objects.ProjectService/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddRouteToProject(ctx context.Context, in *AddRouteToProjectRequest, opts ...grpc.CallOption) (*AddRouteToProjectResponse, error) {
	out := new(AddRouteToProjectResponse)
	err := c.cc.Invoke(ctx, "/objects.ProjectService/AddRouteToProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	AddRouteToProject(context.Context, *AddRouteToProjectRequest) (*AddRouteToProjectResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) AddRouteToProject(context.Context, *AddRouteToProjectRequest) (*AddRouteToProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRouteToProject not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objects.ProjectService/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddRouteToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRouteToProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddRouteToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objects.ProjectService/AddRouteToProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddRouteToProject(ctx, req.(*AddRouteToProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "objects.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "AddRouteToProject",
			Handler:    _ProjectService_AddRouteToProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objects.proto",
}
