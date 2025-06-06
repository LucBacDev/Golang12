// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wallet

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

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetUserByAccountNumber(ctx context.Context, in *GetUserByAccountNumberRequest, opts ...grpc.CallOption) (*GetUserByAccountNumberResponse, error)
	DebitBalance(ctx context.Context, in *DebitRequest, opts ...grpc.CallOption) (*DebitResponse, error)
	CreditBalance(ctx context.Context, in *CreditRequest, opts ...grpc.CallOption) (*CreditResponse, error)
	RefundDebit(ctx context.Context, in *RefundDebitRequest, opts ...grpc.CallOption) (*RefundDebitResponse, error)
	UndoCredit(ctx context.Context, in *UndoCreditRequest, opts ...grpc.CallOption) (*UndoCreditResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetUserByAccountNumber(ctx context.Context, in *GetUserByAccountNumberRequest, opts ...grpc.CallOption) (*GetUserByAccountNumberResponse, error) {
	out := new(GetUserByAccountNumberResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetUserByAccountNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) DebitBalance(ctx context.Context, in *DebitRequest, opts ...grpc.CallOption) (*DebitResponse, error) {
	out := new(DebitResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/DebitBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreditBalance(ctx context.Context, in *CreditRequest, opts ...grpc.CallOption) (*CreditResponse, error) {
	out := new(CreditResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/CreditBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) RefundDebit(ctx context.Context, in *RefundDebitRequest, opts ...grpc.CallOption) (*RefundDebitResponse, error) {
	out := new(RefundDebitResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/RefundDebit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) UndoCredit(ctx context.Context, in *UndoCreditRequest, opts ...grpc.CallOption) (*UndoCreditResponse, error) {
	out := new(UndoCreditResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/UndoCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	GetUserByAccountNumber(context.Context, *GetUserByAccountNumberRequest) (*GetUserByAccountNumberResponse, error)
	DebitBalance(context.Context, *DebitRequest) (*DebitResponse, error)
	CreditBalance(context.Context, *CreditRequest) (*CreditResponse, error)
	RefundDebit(context.Context, *RefundDebitRequest) (*RefundDebitResponse, error)
	UndoCredit(context.Context, *UndoCreditRequest) (*UndoCreditResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) GetUserByAccountNumber(context.Context, *GetUserByAccountNumberRequest) (*GetUserByAccountNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByAccountNumber not implemented")
}
func (UnimplementedWalletServiceServer) DebitBalance(context.Context, *DebitRequest) (*DebitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitBalance not implemented")
}
func (UnimplementedWalletServiceServer) CreditBalance(context.Context, *CreditRequest) (*CreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditBalance not implemented")
}
func (UnimplementedWalletServiceServer) RefundDebit(context.Context, *RefundDebitRequest) (*RefundDebitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundDebit not implemented")
}
func (UnimplementedWalletServiceServer) UndoCredit(context.Context, *UndoCreditRequest) (*UndoCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndoCredit not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetUserByAccountNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByAccountNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetUserByAccountNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetUserByAccountNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetUserByAccountNumber(ctx, req.(*GetUserByAccountNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_DebitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).DebitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/DebitBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).DebitBalance(ctx, req.(*DebitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreditBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreditBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/CreditBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreditBalance(ctx, req.(*CreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_RefundDebit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundDebitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).RefundDebit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/RefundDebit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).RefundDebit(ctx, req.(*RefundDebitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_UndoCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndoCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UndoCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/UndoCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UndoCredit(ctx, req.(*UndoCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByAccountNumber",
			Handler:    _WalletService_GetUserByAccountNumber_Handler,
		},
		{
			MethodName: "DebitBalance",
			Handler:    _WalletService_DebitBalance_Handler,
		},
		{
			MethodName: "CreditBalance",
			Handler:    _WalletService_CreditBalance_Handler,
		},
		{
			MethodName: "RefundDebit",
			Handler:    _WalletService_RefundDebit_Handler,
		},
		{
			MethodName: "UndoCredit",
			Handler:    _WalletService_UndoCredit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang/proto/wallet/wallet.proto",
}
