// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: zitadel/settings/v2/settings_service.proto

package settings

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
	SettingsService_GetGeneralSettings_FullMethodName            = "/zitadel.settings.v2.SettingsService/GetGeneralSettings"
	SettingsService_GetLoginSettings_FullMethodName              = "/zitadel.settings.v2.SettingsService/GetLoginSettings"
	SettingsService_GetActiveIdentityProviders_FullMethodName    = "/zitadel.settings.v2.SettingsService/GetActiveIdentityProviders"
	SettingsService_GetPasswordComplexitySettings_FullMethodName = "/zitadel.settings.v2.SettingsService/GetPasswordComplexitySettings"
	SettingsService_GetPasswordExpirySettings_FullMethodName     = "/zitadel.settings.v2.SettingsService/GetPasswordExpirySettings"
	SettingsService_GetBrandingSettings_FullMethodName           = "/zitadel.settings.v2.SettingsService/GetBrandingSettings"
	SettingsService_GetDomainSettings_FullMethodName             = "/zitadel.settings.v2.SettingsService/GetDomainSettings"
	SettingsService_GetLegalAndSupportSettings_FullMethodName    = "/zitadel.settings.v2.SettingsService/GetLegalAndSupportSettings"
	SettingsService_GetLockoutSettings_FullMethodName            = "/zitadel.settings.v2.SettingsService/GetLockoutSettings"
	SettingsService_GetSecuritySettings_FullMethodName           = "/zitadel.settings.v2.SettingsService/GetSecuritySettings"
	SettingsService_SetSecuritySettings_FullMethodName           = "/zitadel.settings.v2.SettingsService/SetSecuritySettings"
)

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsServiceClient interface {
	// Get basic information over the instance
	GetGeneralSettings(ctx context.Context, in *GetGeneralSettingsRequest, opts ...grpc.CallOption) (*GetGeneralSettingsResponse, error)
	// Get the login settings
	GetLoginSettings(ctx context.Context, in *GetLoginSettingsRequest, opts ...grpc.CallOption) (*GetLoginSettingsResponse, error)
	// Get the current active identity providers
	GetActiveIdentityProviders(ctx context.Context, in *GetActiveIdentityProvidersRequest, opts ...grpc.CallOption) (*GetActiveIdentityProvidersResponse, error)
	// Get the password complexity settings
	GetPasswordComplexitySettings(ctx context.Context, in *GetPasswordComplexitySettingsRequest, opts ...grpc.CallOption) (*GetPasswordComplexitySettingsResponse, error)
	// Get the password expiry settings
	GetPasswordExpirySettings(ctx context.Context, in *GetPasswordExpirySettingsRequest, opts ...grpc.CallOption) (*GetPasswordExpirySettingsResponse, error)
	// Get the current active branding settings
	GetBrandingSettings(ctx context.Context, in *GetBrandingSettingsRequest, opts ...grpc.CallOption) (*GetBrandingSettingsResponse, error)
	// Get the domain settings
	GetDomainSettings(ctx context.Context, in *GetDomainSettingsRequest, opts ...grpc.CallOption) (*GetDomainSettingsResponse, error)
	// Get the legal and support settings
	GetLegalAndSupportSettings(ctx context.Context, in *GetLegalAndSupportSettingsRequest, opts ...grpc.CallOption) (*GetLegalAndSupportSettingsResponse, error)
	// Get the lockout settings
	GetLockoutSettings(ctx context.Context, in *GetLockoutSettingsRequest, opts ...grpc.CallOption) (*GetLockoutSettingsResponse, error)
	// Get the security settings
	GetSecuritySettings(ctx context.Context, in *GetSecuritySettingsRequest, opts ...grpc.CallOption) (*GetSecuritySettingsResponse, error)
	// Set the security settings
	SetSecuritySettings(ctx context.Context, in *SetSecuritySettingsRequest, opts ...grpc.CallOption) (*SetSecuritySettingsResponse, error)
}

type settingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) GetGeneralSettings(ctx context.Context, in *GetGeneralSettingsRequest, opts ...grpc.CallOption) (*GetGeneralSettingsResponse, error) {
	out := new(GetGeneralSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetGeneralSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetLoginSettings(ctx context.Context, in *GetLoginSettingsRequest, opts ...grpc.CallOption) (*GetLoginSettingsResponse, error) {
	out := new(GetLoginSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetLoginSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetActiveIdentityProviders(ctx context.Context, in *GetActiveIdentityProvidersRequest, opts ...grpc.CallOption) (*GetActiveIdentityProvidersResponse, error) {
	out := new(GetActiveIdentityProvidersResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetActiveIdentityProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetPasswordComplexitySettings(ctx context.Context, in *GetPasswordComplexitySettingsRequest, opts ...grpc.CallOption) (*GetPasswordComplexitySettingsResponse, error) {
	out := new(GetPasswordComplexitySettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetPasswordComplexitySettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetPasswordExpirySettings(ctx context.Context, in *GetPasswordExpirySettingsRequest, opts ...grpc.CallOption) (*GetPasswordExpirySettingsResponse, error) {
	out := new(GetPasswordExpirySettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetPasswordExpirySettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetBrandingSettings(ctx context.Context, in *GetBrandingSettingsRequest, opts ...grpc.CallOption) (*GetBrandingSettingsResponse, error) {
	out := new(GetBrandingSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetBrandingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetDomainSettings(ctx context.Context, in *GetDomainSettingsRequest, opts ...grpc.CallOption) (*GetDomainSettingsResponse, error) {
	out := new(GetDomainSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetDomainSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetLegalAndSupportSettings(ctx context.Context, in *GetLegalAndSupportSettingsRequest, opts ...grpc.CallOption) (*GetLegalAndSupportSettingsResponse, error) {
	out := new(GetLegalAndSupportSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetLegalAndSupportSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetLockoutSettings(ctx context.Context, in *GetLockoutSettingsRequest, opts ...grpc.CallOption) (*GetLockoutSettingsResponse, error) {
	out := new(GetLockoutSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetLockoutSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetSecuritySettings(ctx context.Context, in *GetSecuritySettingsRequest, opts ...grpc.CallOption) (*GetSecuritySettingsResponse, error) {
	out := new(GetSecuritySettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetSecuritySettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) SetSecuritySettings(ctx context.Context, in *SetSecuritySettingsRequest, opts ...grpc.CallOption) (*SetSecuritySettingsResponse, error) {
	out := new(SetSecuritySettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_SetSecuritySettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
// All implementations must embed UnimplementedSettingsServiceServer
// for forward compatibility
type SettingsServiceServer interface {
	// Get basic information over the instance
	GetGeneralSettings(context.Context, *GetGeneralSettingsRequest) (*GetGeneralSettingsResponse, error)
	// Get the login settings
	GetLoginSettings(context.Context, *GetLoginSettingsRequest) (*GetLoginSettingsResponse, error)
	// Get the current active identity providers
	GetActiveIdentityProviders(context.Context, *GetActiveIdentityProvidersRequest) (*GetActiveIdentityProvidersResponse, error)
	// Get the password complexity settings
	GetPasswordComplexitySettings(context.Context, *GetPasswordComplexitySettingsRequest) (*GetPasswordComplexitySettingsResponse, error)
	// Get the password expiry settings
	GetPasswordExpirySettings(context.Context, *GetPasswordExpirySettingsRequest) (*GetPasswordExpirySettingsResponse, error)
	// Get the current active branding settings
	GetBrandingSettings(context.Context, *GetBrandingSettingsRequest) (*GetBrandingSettingsResponse, error)
	// Get the domain settings
	GetDomainSettings(context.Context, *GetDomainSettingsRequest) (*GetDomainSettingsResponse, error)
	// Get the legal and support settings
	GetLegalAndSupportSettings(context.Context, *GetLegalAndSupportSettingsRequest) (*GetLegalAndSupportSettingsResponse, error)
	// Get the lockout settings
	GetLockoutSettings(context.Context, *GetLockoutSettingsRequest) (*GetLockoutSettingsResponse, error)
	// Get the security settings
	GetSecuritySettings(context.Context, *GetSecuritySettingsRequest) (*GetSecuritySettingsResponse, error)
	// Set the security settings
	SetSecuritySettings(context.Context, *SetSecuritySettingsRequest) (*SetSecuritySettingsResponse, error)
	mustEmbedUnimplementedSettingsServiceServer()
}

// UnimplementedSettingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServiceServer struct {
}

func (UnimplementedSettingsServiceServer) GetGeneralSettings(context.Context, *GetGeneralSettingsRequest) (*GetGeneralSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeneralSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetLoginSettings(context.Context, *GetLoginSettingsRequest) (*GetLoginSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetActiveIdentityProviders(context.Context, *GetActiveIdentityProvidersRequest) (*GetActiveIdentityProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveIdentityProviders not implemented")
}
func (UnimplementedSettingsServiceServer) GetPasswordComplexitySettings(context.Context, *GetPasswordComplexitySettingsRequest) (*GetPasswordComplexitySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordComplexitySettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetPasswordExpirySettings(context.Context, *GetPasswordExpirySettingsRequest) (*GetPasswordExpirySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordExpirySettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetBrandingSettings(context.Context, *GetBrandingSettingsRequest) (*GetBrandingSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandingSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetDomainSettings(context.Context, *GetDomainSettingsRequest) (*GetDomainSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetLegalAndSupportSettings(context.Context, *GetLegalAndSupportSettingsRequest) (*GetLegalAndSupportSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLegalAndSupportSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetLockoutSettings(context.Context, *GetLockoutSettingsRequest) (*GetLockoutSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLockoutSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetSecuritySettings(context.Context, *GetSecuritySettingsRequest) (*GetSecuritySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecuritySettings not implemented")
}
func (UnimplementedSettingsServiceServer) SetSecuritySettings(context.Context, *SetSecuritySettingsRequest) (*SetSecuritySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSecuritySettings not implemented")
}
func (UnimplementedSettingsServiceServer) mustEmbedUnimplementedSettingsServiceServer() {}

// UnsafeSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServiceServer will
// result in compilation errors.
type UnsafeSettingsServiceServer interface {
	mustEmbedUnimplementedSettingsServiceServer()
}

func RegisterSettingsServiceServer(s grpc.ServiceRegistrar, srv SettingsServiceServer) {
	s.RegisterService(&SettingsService_ServiceDesc, srv)
}

func _SettingsService_GetGeneralSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeneralSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetGeneralSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetGeneralSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetGeneralSettings(ctx, req.(*GetGeneralSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetLoginSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetLoginSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetLoginSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetLoginSettings(ctx, req.(*GetLoginSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetActiveIdentityProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveIdentityProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetActiveIdentityProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetActiveIdentityProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetActiveIdentityProviders(ctx, req.(*GetActiveIdentityProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetPasswordComplexitySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordComplexitySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetPasswordComplexitySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetPasswordComplexitySettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetPasswordComplexitySettings(ctx, req.(*GetPasswordComplexitySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetPasswordExpirySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordExpirySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetPasswordExpirySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetPasswordExpirySettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetPasswordExpirySettings(ctx, req.(*GetPasswordExpirySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetBrandingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetBrandingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetBrandingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetBrandingSettings(ctx, req.(*GetBrandingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetDomainSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetDomainSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetDomainSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetDomainSettings(ctx, req.(*GetDomainSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetLegalAndSupportSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLegalAndSupportSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetLegalAndSupportSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetLegalAndSupportSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetLegalAndSupportSettings(ctx, req.(*GetLegalAndSupportSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetLockoutSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLockoutSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetLockoutSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetLockoutSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetLockoutSettings(ctx, req.(*GetLockoutSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetSecuritySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecuritySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetSecuritySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetSecuritySettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetSecuritySettings(ctx, req.(*GetSecuritySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_SetSecuritySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSecuritySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).SetSecuritySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_SetSecuritySettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).SetSecuritySettings(ctx, req.(*SetSecuritySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingsService_ServiceDesc is the grpc.ServiceDesc for SettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zitadel.settings.v2.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeneralSettings",
			Handler:    _SettingsService_GetGeneralSettings_Handler,
		},
		{
			MethodName: "GetLoginSettings",
			Handler:    _SettingsService_GetLoginSettings_Handler,
		},
		{
			MethodName: "GetActiveIdentityProviders",
			Handler:    _SettingsService_GetActiveIdentityProviders_Handler,
		},
		{
			MethodName: "GetPasswordComplexitySettings",
			Handler:    _SettingsService_GetPasswordComplexitySettings_Handler,
		},
		{
			MethodName: "GetPasswordExpirySettings",
			Handler:    _SettingsService_GetPasswordExpirySettings_Handler,
		},
		{
			MethodName: "GetBrandingSettings",
			Handler:    _SettingsService_GetBrandingSettings_Handler,
		},
		{
			MethodName: "GetDomainSettings",
			Handler:    _SettingsService_GetDomainSettings_Handler,
		},
		{
			MethodName: "GetLegalAndSupportSettings",
			Handler:    _SettingsService_GetLegalAndSupportSettings_Handler,
		},
		{
			MethodName: "GetLockoutSettings",
			Handler:    _SettingsService_GetLockoutSettings_Handler,
		},
		{
			MethodName: "GetSecuritySettings",
			Handler:    _SettingsService_GetSecuritySettings_Handler,
		},
		{
			MethodName: "SetSecuritySettings",
			Handler:    _SettingsService_SetSecuritySettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zitadel/settings/v2/settings_service.proto",
}