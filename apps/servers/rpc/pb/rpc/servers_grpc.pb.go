// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: apps/servers/rpc/servers.proto

package rpc

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

const (
	ServerRpc_GetServers_FullMethodName      = "/serversRPC.ServerRpc/GetServers"
	ServerRpc_GetUniqueServer_FullMethodName = "/serversRPC.ServerRpc/GetUniqueServer"
	ServerRpc_PostServer_FullMethodName      = "/serversRPC.ServerRpc/PostServer"
	ServerRpc_PatchServer_FullMethodName     = "/serversRPC.ServerRpc/PatchServer"
	ServerRpc_DeleteServer_FullMethodName    = "/serversRPC.ServerRpc/DeleteServer"
)

// ServerRpcClient is the client API for ServerRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerRpcClient interface {
	GetServers(ctx context.Context, in *GetServersReq, opts ...grpc.CallOption) (*GetServersRes, error)
	GetUniqueServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerRes, error)
	PostServer(ctx context.Context, in *PostServerReq, opts ...grpc.CallOption) (*EmptyRes, error)
	PatchServer(ctx context.Context, in *PatchServerReq, opts ...grpc.CallOption) (*EmptyRes, error)
	DeleteServer(ctx context.Context, in *DeleteServerReq, opts ...grpc.CallOption) (*EmptyRes, error)
}

type serverRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewServerRpcClient(cc grpc.ClientConnInterface) ServerRpcClient {
	return &serverRpcClient{cc}
}

func (c *serverRpcClient) GetServers(ctx context.Context, in *GetServersReq, opts ...grpc.CallOption) (*GetServersRes, error) {
	out := new(GetServersRes)
	err := c.cc.Invoke(ctx, ServerRpc_GetServers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverRpcClient) GetUniqueServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerRes, error) {
	out := new(GetServerRes)
	err := c.cc.Invoke(ctx, ServerRpc_GetUniqueServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverRpcClient) PostServer(ctx context.Context, in *PostServerReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, ServerRpc_PostServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverRpcClient) PatchServer(ctx context.Context, in *PatchServerReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, ServerRpc_PatchServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverRpcClient) DeleteServer(ctx context.Context, in *DeleteServerReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, ServerRpc_DeleteServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerRpcServer is the server API for ServerRpc service.
// All implementations must embed UnimplementedServerRpcServer
// for forward compatibility
type ServerRpcServer interface {
	GetServers(context.Context, *GetServersReq) (*GetServersRes, error)
	GetUniqueServer(context.Context, *GetServerReq) (*GetServerRes, error)
	PostServer(context.Context, *PostServerReq) (*EmptyRes, error)
	PatchServer(context.Context, *PatchServerReq) (*EmptyRes, error)
	DeleteServer(context.Context, *DeleteServerReq) (*EmptyRes, error)
	mustEmbedUnimplementedServerRpcServer()
}

// UnimplementedServerRpcServer must be embedded to have forward compatible implementations.
type UnimplementedServerRpcServer struct {
}

func (UnimplementedServerRpcServer) GetServers(context.Context, *GetServersReq) (*GetServersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (UnimplementedServerRpcServer) GetUniqueServer(context.Context, *GetServerReq) (*GetServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUniqueServer not implemented")
}
func (UnimplementedServerRpcServer) PostServer(context.Context, *PostServerReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostServer not implemented")
}
func (UnimplementedServerRpcServer) PatchServer(context.Context, *PatchServerReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchServer not implemented")
}
func (UnimplementedServerRpcServer) DeleteServer(context.Context, *DeleteServerReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedServerRpcServer) mustEmbedUnimplementedServerRpcServer() {}

// UnsafeServerRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerRpcServer will
// result in compilation errors.
type UnsafeServerRpcServer interface {
	mustEmbedUnimplementedServerRpcServer()
}

func RegisterServerRpcServer(s grpc.ServiceRegistrar, srv ServerRpcServer) {
	s.RegisterService(&ServerRpc_ServiceDesc, srv)
}

func _ServerRpc_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRpcServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRpc_GetServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRpcServer).GetServers(ctx, req.(*GetServersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerRpc_GetUniqueServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRpcServer).GetUniqueServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRpc_GetUniqueServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRpcServer).GetUniqueServer(ctx, req.(*GetServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerRpc_PostServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRpcServer).PostServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRpc_PostServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRpcServer).PostServer(ctx, req.(*PostServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerRpc_PatchServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRpcServer).PatchServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRpc_PatchServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRpcServer).PatchServer(ctx, req.(*PatchServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerRpc_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRpcServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRpc_DeleteServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRpcServer).DeleteServer(ctx, req.(*DeleteServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerRpc_ServiceDesc is the grpc.ServiceDesc for ServerRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serversRPC.ServerRpc",
	HandlerType: (*ServerRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServers",
			Handler:    _ServerRpc_GetServers_Handler,
		},
		{
			MethodName: "GetUniqueServer",
			Handler:    _ServerRpc_GetUniqueServer_Handler,
		},
		{
			MethodName: "PostServer",
			Handler:    _ServerRpc_PostServer_Handler,
		},
		{
			MethodName: "PatchServer",
			Handler:    _ServerRpc_PatchServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _ServerRpc_DeleteServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/servers/rpc/servers.proto",
}

const (
	MemberRpc_InviteMember_FullMethodName = "/serversRPC.MemberRpc/InviteMember"
	MemberRpc_PatchMember_FullMethodName  = "/serversRPC.MemberRpc/PatchMember"
	MemberRpc_DeleteMember_FullMethodName = "/serversRPC.MemberRpc/DeleteMember"
)

// MemberRpcClient is the client API for MemberRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberRpcClient interface {
	InviteMember(ctx context.Context, in *InviteMemberReq, opts ...grpc.CallOption) (*GetServerRes, error)
	PatchMember(ctx context.Context, in *PatchMemberReq, opts ...grpc.CallOption) (*GetServerRes, error)
	DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*GetServerRes, error)
}

type memberRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberRpcClient(cc grpc.ClientConnInterface) MemberRpcClient {
	return &memberRpcClient{cc}
}

func (c *memberRpcClient) InviteMember(ctx context.Context, in *InviteMemberReq, opts ...grpc.CallOption) (*GetServerRes, error) {
	out := new(GetServerRes)
	err := c.cc.Invoke(ctx, MemberRpc_InviteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberRpcClient) PatchMember(ctx context.Context, in *PatchMemberReq, opts ...grpc.CallOption) (*GetServerRes, error) {
	out := new(GetServerRes)
	err := c.cc.Invoke(ctx, MemberRpc_PatchMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberRpcClient) DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*GetServerRes, error) {
	out := new(GetServerRes)
	err := c.cc.Invoke(ctx, MemberRpc_DeleteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberRpcServer is the server API for MemberRpc service.
// All implementations must embed UnimplementedMemberRpcServer
// for forward compatibility
type MemberRpcServer interface {
	InviteMember(context.Context, *InviteMemberReq) (*GetServerRes, error)
	PatchMember(context.Context, *PatchMemberReq) (*GetServerRes, error)
	DeleteMember(context.Context, *DeleteMemberReq) (*GetServerRes, error)
	mustEmbedUnimplementedMemberRpcServer()
}

// UnimplementedMemberRpcServer must be embedded to have forward compatible implementations.
type UnimplementedMemberRpcServer struct {
}

func (UnimplementedMemberRpcServer) InviteMember(context.Context, *InviteMemberReq) (*GetServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteMember not implemented")
}
func (UnimplementedMemberRpcServer) PatchMember(context.Context, *PatchMemberReq) (*GetServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchMember not implemented")
}
func (UnimplementedMemberRpcServer) DeleteMember(context.Context, *DeleteMemberReq) (*GetServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedMemberRpcServer) mustEmbedUnimplementedMemberRpcServer() {}

// UnsafeMemberRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberRpcServer will
// result in compilation errors.
type UnsafeMemberRpcServer interface {
	mustEmbedUnimplementedMemberRpcServer()
}

func RegisterMemberRpcServer(s grpc.ServiceRegistrar, srv MemberRpcServer) {
	s.RegisterService(&MemberRpc_ServiceDesc, srv)
}

func _MemberRpc_InviteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberRpcServer).InviteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberRpc_InviteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberRpcServer).InviteMember(ctx, req.(*InviteMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberRpc_PatchMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberRpcServer).PatchMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberRpc_PatchMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberRpcServer).PatchMember(ctx, req.(*PatchMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberRpc_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberRpcServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberRpc_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberRpcServer).DeleteMember(ctx, req.(*DeleteMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberRpc_ServiceDesc is the grpc.ServiceDesc for MemberRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serversRPC.MemberRpc",
	HandlerType: (*MemberRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteMember",
			Handler:    _MemberRpc_InviteMember_Handler,
		},
		{
			MethodName: "PatchMember",
			Handler:    _MemberRpc_PatchMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _MemberRpc_DeleteMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/servers/rpc/servers.proto",
}
