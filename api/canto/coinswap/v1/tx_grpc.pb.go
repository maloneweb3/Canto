// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: canto/coinswap/v1/tx.proto

package coinswapv1

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
	Msg_AddLiquidity_FullMethodName    = "/canto.coinswap.v1.Msg/AddLiquidity"
	Msg_RemoveLiquidity_FullMethodName = "/canto.coinswap.v1.Msg/RemoveLiquidity"
	Msg_SwapCoin_FullMethodName        = "/canto.coinswap.v1.Msg/SwapCoin"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// AddLiquidity defines a method for depositing some tokens to the liquidity
	// pool
	AddLiquidity(ctx context.Context, in *MsgAddLiquidity, opts ...grpc.CallOption) (*MsgAddLiquidityResponse, error)
	// RemoveLiquidity defines a method for withdraw some tokens from the
	// liquidity pool
	RemoveLiquidity(ctx context.Context, in *MsgRemoveLiquidity, opts ...grpc.CallOption) (*MsgRemoveLiquidityResponse, error)
	// SwapCoin defines a method for swapping a token with the other token from
	// the liquidity pool
	SwapCoin(ctx context.Context, in *MsgSwapOrder, opts ...grpc.CallOption) (*MsgSwapCoinResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddLiquidity(ctx context.Context, in *MsgAddLiquidity, opts ...grpc.CallOption) (*MsgAddLiquidityResponse, error) {
	out := new(MsgAddLiquidityResponse)
	err := c.cc.Invoke(ctx, Msg_AddLiquidity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveLiquidity(ctx context.Context, in *MsgRemoveLiquidity, opts ...grpc.CallOption) (*MsgRemoveLiquidityResponse, error) {
	out := new(MsgRemoveLiquidityResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveLiquidity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapCoin(ctx context.Context, in *MsgSwapOrder, opts ...grpc.CallOption) (*MsgSwapCoinResponse, error) {
	out := new(MsgSwapCoinResponse)
	err := c.cc.Invoke(ctx, Msg_SwapCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// AddLiquidity defines a method for depositing some tokens to the liquidity
	// pool
	AddLiquidity(context.Context, *MsgAddLiquidity) (*MsgAddLiquidityResponse, error)
	// RemoveLiquidity defines a method for withdraw some tokens from the
	// liquidity pool
	RemoveLiquidity(context.Context, *MsgRemoveLiquidity) (*MsgRemoveLiquidityResponse, error)
	// SwapCoin defines a method for swapping a token with the other token from
	// the liquidity pool
	SwapCoin(context.Context, *MsgSwapOrder) (*MsgSwapCoinResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) AddLiquidity(context.Context, *MsgAddLiquidity) (*MsgAddLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLiquidity not implemented")
}
func (UnimplementedMsgServer) RemoveLiquidity(context.Context, *MsgRemoveLiquidity) (*MsgRemoveLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLiquidity not implemented")
}
func (UnimplementedMsgServer) SwapCoin(context.Context, *MsgSwapOrder) (*MsgSwapCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapCoin not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AddLiquidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddLiquidity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddLiquidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddLiquidity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddLiquidity(ctx, req.(*MsgAddLiquidity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveLiquidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveLiquidity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveLiquidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveLiquidity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveLiquidity(ctx, req.(*MsgRemoveLiquidity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SwapCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapCoin(ctx, req.(*MsgSwapOrder))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "canto.coinswap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLiquidity",
			Handler:    _Msg_AddLiquidity_Handler,
		},
		{
			MethodName: "RemoveLiquidity",
			Handler:    _Msg_RemoveLiquidity_Handler,
		},
		{
			MethodName: "SwapCoin",
			Handler:    _Msg_SwapCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canto/coinswap/v1/tx.proto",
}
