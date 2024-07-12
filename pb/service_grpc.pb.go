// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Service_ListTransactions_FullMethodName  = "/transaction.Service/ListTransactions"
	Service_CreateTransaction_FullMethodName = "/transaction.Service/CreateTransaction"
	Service_UpdateTransaction_FullMethodName = "/transaction.Service/UpdateTransaction"
	Service_DeleteTransaction_FullMethodName = "/transaction.Service/DeleteTransaction"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, Service_ListTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, Service_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTransactionResponse)
	err := c.cc.Invoke(ctx, Service_UpdateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTransactionResponse)
	err := c.cc.Invoke(ctx, Service_DeleteTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations should embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error)
}

// UnimplementedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
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

func _Service_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransactions",
			Handler:    _Service_ListTransactions_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Service_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _Service_UpdateTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _Service_DeleteTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
