// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: dish.proto

package orderproto

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
	DishService_GetOrderDish_FullMethodName = "/order.DishService/GetOrderDish"
)

// DishServiceClient is the client API for DishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DishServiceClient interface {
	GetOrderDish(ctx context.Context, in *GetOrderDishRequest, opts ...grpc.CallOption) (*GetOrderDishResponse, error)
}

type dishServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDishServiceClient(cc grpc.ClientConnInterface) DishServiceClient {
	return &dishServiceClient{cc}
}

func (c *dishServiceClient) GetOrderDish(ctx context.Context, in *GetOrderDishRequest, opts ...grpc.CallOption) (*GetOrderDishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderDishResponse)
	err := c.cc.Invoke(ctx, DishService_GetOrderDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DishServiceServer is the server API for DishService service.
// All implementations must embed UnimplementedDishServiceServer
// for forward compatibility
type DishServiceServer interface {
	GetOrderDish(context.Context, *GetOrderDishRequest) (*GetOrderDishResponse, error)
	mustEmbedUnimplementedDishServiceServer()
}

// UnimplementedDishServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDishServiceServer struct {
}

func (UnimplementedDishServiceServer) GetOrderDish(context.Context, *GetOrderDishRequest) (*GetOrderDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDish not implemented")
}
func (UnimplementedDishServiceServer) mustEmbedUnimplementedDishServiceServer() {}

// UnsafeDishServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DishServiceServer will
// result in compilation errors.
type UnsafeDishServiceServer interface {
	mustEmbedUnimplementedDishServiceServer()
}

func RegisterDishServiceServer(s grpc.ServiceRegistrar, srv DishServiceServer) {
	s.RegisterService(&DishService_ServiceDesc, srv)
}

func _DishService_GetOrderDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServiceServer).GetOrderDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DishService_GetOrderDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServiceServer).GetOrderDish(ctx, req.(*GetOrderDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DishService_ServiceDesc is the grpc.ServiceDesc for DishService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DishService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.DishService",
	HandlerType: (*DishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderDish",
			Handler:    _DishService_GetOrderDish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dish.proto",
}
