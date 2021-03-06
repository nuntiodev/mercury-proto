// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_mercury

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Ping(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	CreateUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	GetUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	GetAllUsers(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	UpdateUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	DeleteUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	CreateConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	AddUserToConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	RemoveUserFromConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	Send(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	ListConversations(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	DeleteConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
	Heartbeat(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Ping(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetAllUsers(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteUser(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/CreateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddUserToConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/AddUserToConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RemoveUserFromConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/RemoveUserFromConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Send(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListConversations(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/ListConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteConversation(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/DeleteConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Heartbeat(ctx context.Context, in *MercuryRequest, opts ...grpc.CallOption) (*MercuryResponse, error) {
	out := new(MercuryResponse)
	err := c.cc.Invoke(ctx, "/Mercury.Service/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations should embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Ping(context.Context, *MercuryRequest) (*MercuryResponse, error)
	CreateUser(context.Context, *MercuryRequest) (*MercuryResponse, error)
	GetUser(context.Context, *MercuryRequest) (*MercuryResponse, error)
	GetAllUsers(context.Context, *MercuryRequest) (*MercuryResponse, error)
	UpdateUser(context.Context, *MercuryRequest) (*MercuryResponse, error)
	DeleteUser(context.Context, *MercuryRequest) (*MercuryResponse, error)
	CreateConversation(context.Context, *MercuryRequest) (*MercuryResponse, error)
	AddUserToConversation(context.Context, *MercuryRequest) (*MercuryResponse, error)
	RemoveUserFromConversation(context.Context, *MercuryRequest) (*MercuryResponse, error)
	Send(context.Context, *MercuryRequest) (*MercuryResponse, error)
	ListConversations(context.Context, *MercuryRequest) (*MercuryResponse, error)
	DeleteConversation(context.Context, *MercuryRequest) (*MercuryResponse, error)
	Heartbeat(context.Context, *MercuryRequest) (*MercuryResponse, error)
}

// UnimplementedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Ping(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedServiceServer) CreateUser(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedServiceServer) GetUser(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedServiceServer) GetAllUsers(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedServiceServer) UpdateUser(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedServiceServer) DeleteUser(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedServiceServer) CreateConversation(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConversation not implemented")
}
func (UnimplementedServiceServer) AddUserToConversation(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToConversation not implemented")
}
func (UnimplementedServiceServer) RemoveUserFromConversation(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromConversation not implemented")
}
func (UnimplementedServiceServer) Send(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedServiceServer) ListConversations(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConversations not implemented")
}
func (UnimplementedServiceServer) DeleteConversation(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConversation not implemented")
}
func (UnimplementedServiceServer) Heartbeat(context.Context, *MercuryRequest) (*MercuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Ping(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateUser(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUser(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetAllUsers(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateUser(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteUser(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateConversation(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddUserToConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddUserToConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/AddUserToConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddUserToConversation(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RemoveUserFromConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RemoveUserFromConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/RemoveUserFromConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RemoveUserFromConversation(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Send(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/ListConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListConversations(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/DeleteConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteConversation(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mercury.Service/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Heartbeat(ctx, req.(*MercuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mercury.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Service_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Service_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Service_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _Service_GetAllUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Service_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Service_DeleteUser_Handler,
		},
		{
			MethodName: "CreateConversation",
			Handler:    _Service_CreateConversation_Handler,
		},
		{
			MethodName: "AddUserToConversation",
			Handler:    _Service_AddUserToConversation_Handler,
		},
		{
			MethodName: "RemoveUserFromConversation",
			Handler:    _Service_RemoveUserFromConversation_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Service_Send_Handler,
		},
		{
			MethodName: "ListConversations",
			Handler:    _Service_ListConversations_Handler,
		},
		{
			MethodName: "DeleteConversation",
			Handler:    _Service_DeleteConversation_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Service_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}
