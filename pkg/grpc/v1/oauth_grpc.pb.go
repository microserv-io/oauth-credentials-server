// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.3
// source: pkg/grpc/v1/oauth.proto

package grpcoauthv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthServiceClient interface {
	ListOAuths(ctx context.Context, in *ListOAuthsRequest, opts ...grpc.CallOption) (OAuthService_ListOAuthsClient, error)
	GetOAuthByID(ctx context.Context, in *GetOAuthByIDRequest, opts ...grpc.CallOption) (*OAuth, error)
	GetOAuthByProvider(ctx context.Context, in *GetOAuthByProviderRequest, opts ...grpc.CallOption) (*OAuth, error)
	GetOAuthCredentialsByProvider(ctx context.Context, in *GetOAuthByProviderRequest, opts ...grpc.CallOption) (*OAuthCredential, error)
}

type oAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthServiceClient(cc grpc.ClientConnInterface) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) ListOAuths(ctx context.Context, in *ListOAuthsRequest, opts ...grpc.CallOption) (OAuthService_ListOAuthsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OAuthService_ServiceDesc.Streams[0], "/grpcoauthv1.OAuthService/ListOAuths", opts...)
	if err != nil {
		return nil, err
	}
	x := &oAuthServiceListOAuthsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OAuthService_ListOAuthsClient interface {
	Recv() (*OAuth, error)
	grpc.ClientStream
}

type oAuthServiceListOAuthsClient struct {
	grpc.ClientStream
}

func (x *oAuthServiceListOAuthsClient) Recv() (*OAuth, error) {
	m := new(OAuth)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oAuthServiceClient) GetOAuthByID(ctx context.Context, in *GetOAuthByIDRequest, opts ...grpc.CallOption) (*OAuth, error) {
	out := new(OAuth)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthService/GetOAuthByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) GetOAuthByProvider(ctx context.Context, in *GetOAuthByProviderRequest, opts ...grpc.CallOption) (*OAuth, error) {
	out := new(OAuth)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthService/GetOAuthByProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) GetOAuthCredentialsByProvider(ctx context.Context, in *GetOAuthByProviderRequest, opts ...grpc.CallOption) (*OAuthCredential, error) {
	out := new(OAuthCredential)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthService/GetOAuthCredentialsByProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
// All implementations must embed UnimplementedOAuthServiceServer
// for forward compatibility
type OAuthServiceServer interface {
	ListOAuths(*ListOAuthsRequest, OAuthService_ListOAuthsServer) error
	GetOAuthByID(context.Context, *GetOAuthByIDRequest) (*OAuth, error)
	GetOAuthByProvider(context.Context, *GetOAuthByProviderRequest) (*OAuth, error)
	GetOAuthCredentialsByProvider(context.Context, *GetOAuthByProviderRequest) (*OAuthCredential, error)
	mustEmbedUnimplementedOAuthServiceServer()
}

// UnimplementedOAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (UnimplementedOAuthServiceServer) ListOAuths(*ListOAuthsRequest, OAuthService_ListOAuthsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOAuths not implemented")
}
func (UnimplementedOAuthServiceServer) GetOAuthByID(context.Context, *GetOAuthByIDRequest) (*OAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuthByID not implemented")
}
func (UnimplementedOAuthServiceServer) GetOAuthByProvider(context.Context, *GetOAuthByProviderRequest) (*OAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuthByProvider not implemented")
}
func (UnimplementedOAuthServiceServer) GetOAuthCredentialsByProvider(context.Context, *GetOAuthByProviderRequest) (*OAuthCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuthCredentialsByProvider not implemented")
}
func (UnimplementedOAuthServiceServer) mustEmbedUnimplementedOAuthServiceServer() {}

// UnsafeOAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthServiceServer will
// result in compilation errors.
type UnsafeOAuthServiceServer interface {
	mustEmbedUnimplementedOAuthServiceServer()
}

func RegisterOAuthServiceServer(s grpc.ServiceRegistrar, srv OAuthServiceServer) {
	s.RegisterService(&OAuthService_ServiceDesc, srv)
}

func _OAuthService_ListOAuths_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOAuthsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OAuthServiceServer).ListOAuths(m, &oAuthServiceListOAuthsServer{stream})
}

type OAuthService_ListOAuthsServer interface {
	Send(*OAuth) error
	grpc.ServerStream
}

type oAuthServiceListOAuthsServer struct {
	grpc.ServerStream
}

func (x *oAuthServiceListOAuthsServer) Send(m *OAuth) error {
	return x.ServerStream.SendMsg(m)
}

func _OAuthService_GetOAuthByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuthByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).GetOAuthByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthService/GetOAuthByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).GetOAuthByID(ctx, req.(*GetOAuthByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_GetOAuthByProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuthByProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).GetOAuthByProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthService/GetOAuthByProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).GetOAuthByProvider(ctx, req.(*GetOAuthByProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_GetOAuthCredentialsByProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuthByProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).GetOAuthCredentialsByProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthService/GetOAuthCredentialsByProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).GetOAuthCredentialsByProvider(ctx, req.(*GetOAuthByProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthService_ServiceDesc is the grpc.ServiceDesc for OAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcoauthv1.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOAuthByID",
			Handler:    _OAuthService_GetOAuthByID_Handler,
		},
		{
			MethodName: "GetOAuthByProvider",
			Handler:    _OAuthService_GetOAuthByProvider_Handler,
		},
		{
			MethodName: "GetOAuthCredentialsByProvider",
			Handler:    _OAuthService_GetOAuthCredentialsByProvider_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOAuths",
			Handler:       _OAuthService_ListOAuths_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/v1/oauth.proto",
}

// OAuthProviderServiceClient is the client API for OAuthProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthProviderServiceClient interface {
	ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OAuthProviderService_ListProvidersClient, error)
	CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*OAuthProvider, error)
	UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*OAuthProvider, error)
	DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type oAuthProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthProviderServiceClient(cc grpc.ClientConnInterface) OAuthProviderServiceClient {
	return &oAuthProviderServiceClient{cc}
}

func (c *oAuthProviderServiceClient) ListProviders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OAuthProviderService_ListProvidersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OAuthProviderService_ServiceDesc.Streams[0], "/grpcoauthv1.OAuthProviderService/ListProviders", opts...)
	if err != nil {
		return nil, err
	}
	x := &oAuthProviderServiceListProvidersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OAuthProviderService_ListProvidersClient interface {
	Recv() (*OAuthProvider, error)
	grpc.ClientStream
}

type oAuthProviderServiceListProvidersClient struct {
	grpc.ClientStream
}

func (x *oAuthProviderServiceListProvidersClient) Recv() (*OAuthProvider, error) {
	m := new(OAuthProvider)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oAuthProviderServiceClient) CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*OAuthProvider, error) {
	out := new(OAuthProvider)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthProviderService/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthProviderServiceClient) UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*OAuthProvider, error) {
	out := new(OAuthProvider)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthProviderService/UpdateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthProviderServiceClient) DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/grpcoauthv1.OAuthProviderService/DeleteProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthProviderServiceServer is the server API for OAuthProviderService service.
// All implementations must embed UnimplementedOAuthProviderServiceServer
// for forward compatibility
type OAuthProviderServiceServer interface {
	ListProviders(*emptypb.Empty, OAuthProviderService_ListProvidersServer) error
	CreateProvider(context.Context, *CreateProviderRequest) (*OAuthProvider, error)
	UpdateProvider(context.Context, *UpdateProviderRequest) (*OAuthProvider, error)
	DeleteProvider(context.Context, *DeleteProviderRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOAuthProviderServiceServer()
}

// UnimplementedOAuthProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthProviderServiceServer struct {
}

func (UnimplementedOAuthProviderServiceServer) ListProviders(*emptypb.Empty, OAuthProviderService_ListProvidersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedOAuthProviderServiceServer) CreateProvider(context.Context, *CreateProviderRequest) (*OAuthProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedOAuthProviderServiceServer) UpdateProvider(context.Context, *UpdateProviderRequest) (*OAuthProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (UnimplementedOAuthProviderServiceServer) DeleteProvider(context.Context, *DeleteProviderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
func (UnimplementedOAuthProviderServiceServer) mustEmbedUnimplementedOAuthProviderServiceServer() {}

// UnsafeOAuthProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthProviderServiceServer will
// result in compilation errors.
type UnsafeOAuthProviderServiceServer interface {
	mustEmbedUnimplementedOAuthProviderServiceServer()
}

func RegisterOAuthProviderServiceServer(s grpc.ServiceRegistrar, srv OAuthProviderServiceServer) {
	s.RegisterService(&OAuthProviderService_ServiceDesc, srv)
}

func _OAuthProviderService_ListProviders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OAuthProviderServiceServer).ListProviders(m, &oAuthProviderServiceListProvidersServer{stream})
}

type OAuthProviderService_ListProvidersServer interface {
	Send(*OAuthProvider) error
	grpc.ServerStream
}

type oAuthProviderServiceListProvidersServer struct {
	grpc.ServerStream
}

func (x *oAuthProviderServiceListProvidersServer) Send(m *OAuthProvider) error {
	return x.ServerStream.SendMsg(m)
}

func _OAuthProviderService_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthProviderServiceServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthProviderService/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthProviderServiceServer).CreateProvider(ctx, req.(*CreateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthProviderService_UpdateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthProviderServiceServer).UpdateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthProviderService/UpdateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthProviderServiceServer).UpdateProvider(ctx, req.(*UpdateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthProviderService_DeleteProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthProviderServiceServer).DeleteProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcoauthv1.OAuthProviderService/DeleteProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthProviderServiceServer).DeleteProvider(ctx, req.(*DeleteProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthProviderService_ServiceDesc is the grpc.ServiceDesc for OAuthProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcoauthv1.OAuthProviderService",
	HandlerType: (*OAuthProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProvider",
			Handler:    _OAuthProviderService_CreateProvider_Handler,
		},
		{
			MethodName: "UpdateProvider",
			Handler:    _OAuthProviderService_UpdateProvider_Handler,
		},
		{
			MethodName: "DeleteProvider",
			Handler:    _OAuthProviderService_DeleteProvider_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProviders",
			Handler:       _OAuthProviderService_ListProviders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/v1/oauth.proto",
}
