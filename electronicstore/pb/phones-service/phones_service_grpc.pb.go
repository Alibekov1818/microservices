// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: phones_service.proto

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

const (
	PhonesService_GetPhone_FullMethodName    = "/auth.PhonesService/GetPhone"
	PhonesService_GetPhones_FullMethodName   = "/auth.PhonesService/GetPhones"
	PhonesService_CreatePhone_FullMethodName = "/auth.PhonesService/CreatePhone"
	PhonesService_DeletePhone_FullMethodName = "/auth.PhonesService/DeletePhone"
	PhonesService_UpdatePhone_FullMethodName = "/auth.PhonesService/UpdatePhone"
)

// PhonesServiceClient is the client API for PhonesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhonesServiceClient interface {
	GetPhone(ctx context.Context, in *PhoneId, opts ...grpc.CallOption) (*Phone, error)
	GetPhones(ctx context.Context, in *GetPhonesRequest, opts ...grpc.CallOption) (*PhoneList, error)
	CreatePhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Phone, error)
	DeletePhone(ctx context.Context, in *PhoneId, opts ...grpc.CallOption) (*Phone, error)
	UpdatePhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Phone, error)
}

type phonesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhonesServiceClient(cc grpc.ClientConnInterface) PhonesServiceClient {
	return &phonesServiceClient{cc}
}

func (c *phonesServiceClient) GetPhone(ctx context.Context, in *PhoneId, opts ...grpc.CallOption) (*Phone, error) {
	out := new(Phone)
	err := c.cc.Invoke(ctx, PhonesService_GetPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phonesServiceClient) GetPhones(ctx context.Context, in *GetPhonesRequest, opts ...grpc.CallOption) (*PhoneList, error) {
	out := new(PhoneList)
	err := c.cc.Invoke(ctx, PhonesService_GetPhones_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phonesServiceClient) CreatePhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Phone, error) {
	out := new(Phone)
	err := c.cc.Invoke(ctx, PhonesService_CreatePhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phonesServiceClient) DeletePhone(ctx context.Context, in *PhoneId, opts ...grpc.CallOption) (*Phone, error) {
	out := new(Phone)
	err := c.cc.Invoke(ctx, PhonesService_DeletePhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phonesServiceClient) UpdatePhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Phone, error) {
	out := new(Phone)
	err := c.cc.Invoke(ctx, PhonesService_UpdatePhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhonesServiceServer is the server API for PhonesService service.
// All implementations must embed UnimplementedPhonesServiceServer
// for forward compatibility
type PhonesServiceServer interface {
	GetPhone(context.Context, *PhoneId) (*Phone, error)
	GetPhones(context.Context, *GetPhonesRequest) (*PhoneList, error)
	CreatePhone(context.Context, *Phone) (*Phone, error)
	DeletePhone(context.Context, *PhoneId) (*Phone, error)
	UpdatePhone(context.Context, *Phone) (*Phone, error)
	mustEmbedUnimplementedPhonesServiceServer()
}

// UnimplementedPhonesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPhonesServiceServer struct {
}

func (UnimplementedPhonesServiceServer) GetPhone(context.Context, *PhoneId) (*Phone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhone not implemented")
}
func (UnimplementedPhonesServiceServer) GetPhones(context.Context, *GetPhonesRequest) (*PhoneList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhones not implemented")
}
func (UnimplementedPhonesServiceServer) CreatePhone(context.Context, *Phone) (*Phone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePhone not implemented")
}
func (UnimplementedPhonesServiceServer) DeletePhone(context.Context, *PhoneId) (*Phone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhone not implemented")
}
func (UnimplementedPhonesServiceServer) UpdatePhone(context.Context, *Phone) (*Phone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhone not implemented")
}
func (UnimplementedPhonesServiceServer) mustEmbedUnimplementedPhonesServiceServer() {}

// UnsafePhonesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhonesServiceServer will
// result in compilation errors.
type UnsafePhonesServiceServer interface {
	mustEmbedUnimplementedPhonesServiceServer()
}

func RegisterPhonesServiceServer(s grpc.ServiceRegistrar, srv PhonesServiceServer) {
	s.RegisterService(&PhonesService_ServiceDesc, srv)
}

func _PhonesService_GetPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServiceServer).GetPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhonesService_GetPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServiceServer).GetPhone(ctx, req.(*PhoneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhonesService_GetPhones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServiceServer).GetPhones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhonesService_GetPhones_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServiceServer).GetPhones(ctx, req.(*GetPhonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhonesService_CreatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServiceServer).CreatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhonesService_CreatePhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServiceServer).CreatePhone(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhonesService_DeletePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServiceServer).DeletePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhonesService_DeletePhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServiceServer).DeletePhone(ctx, req.(*PhoneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhonesService_UpdatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServiceServer).UpdatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhonesService_UpdatePhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServiceServer).UpdatePhone(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

// PhonesService_ServiceDesc is the grpc.ServiceDesc for PhonesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhonesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.PhonesService",
	HandlerType: (*PhonesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhone",
			Handler:    _PhonesService_GetPhone_Handler,
		},
		{
			MethodName: "GetPhones",
			Handler:    _PhonesService_GetPhones_Handler,
		},
		{
			MethodName: "CreatePhone",
			Handler:    _PhonesService_CreatePhone_Handler,
		},
		{
			MethodName: "DeletePhone",
			Handler:    _PhonesService_DeletePhone_Handler,
		},
		{
			MethodName: "UpdatePhone",
			Handler:    _PhonesService_UpdatePhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phones_service.proto",
}
