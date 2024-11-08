// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: src/proto/orderingsystem.proto

package orderingsystem

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

// OrderManagementServiceClient is the client API for OrderManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementServiceClient interface {
	GetOrder(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	SearchOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (OrderManagementService_SearchOrdersClient, error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagementService_UpdateOrdersClient, error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagementService_ProcessOrdersClient, error)
}

type orderManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementServiceClient(cc grpc.ClientConnInterface) OrderManagementServiceClient {
	return &orderManagementServiceClient{cc}
}

func (c *orderManagementServiceClient) GetOrder(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/orderingsystem.OrderManagementService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementServiceClient) SearchOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (OrderManagementService_SearchOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagementService_ServiceDesc.Streams[0], "/orderingsystem.OrderManagementService/SearchOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementServiceSearchOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagementService_SearchOrdersClient interface {
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderManagementServiceSearchOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementServiceSearchOrdersClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementServiceClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagementService_UpdateOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagementService_ServiceDesc.Streams[1], "/orderingsystem.OrderManagementService/UpdateOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementServiceUpdateOrdersClient{stream}
	return x, nil
}

type OrderManagementService_UpdateOrdersClient interface {
	Send(*OrdersRequest) error
	CloseAndRecv() (*OrdersResponse, error)
	grpc.ClientStream
}

type orderManagementServiceUpdateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementServiceUpdateOrdersClient) Send(m *OrdersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementServiceUpdateOrdersClient) CloseAndRecv() (*OrdersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OrdersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementServiceClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagementService_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagementService_ServiceDesc.Streams[2], "/orderingsystem.OrderManagementService/ProcessOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementServiceProcessOrdersClient{stream}
	return x, nil
}

type OrderManagementService_ProcessOrdersClient interface {
	Send(*OrdersRequest) error
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderManagementServiceProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementServiceProcessOrdersClient) Send(m *OrdersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementServiceProcessOrdersClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServiceServer is the server API for OrderManagementService service.
// All implementations must embed UnimplementedOrderManagementServiceServer
// for forward compatibility
type OrderManagementServiceServer interface {
	GetOrder(context.Context, *OrdersRequest) (*OrdersResponse, error)
	SearchOrders(*OrdersRequest, OrderManagementService_SearchOrdersServer) error
	UpdateOrders(OrderManagementService_UpdateOrdersServer) error
	ProcessOrders(OrderManagementService_ProcessOrdersServer) error
	mustEmbedUnimplementedOrderManagementServiceServer()
}

// UnimplementedOrderManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServiceServer struct {
}

func (UnimplementedOrderManagementServiceServer) GetOrder(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) SearchOrders(*OrdersRequest, OrderManagementService_SearchOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrders not implemented")
}
func (UnimplementedOrderManagementServiceServer) UpdateOrders(OrderManagementService_UpdateOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedOrderManagementServiceServer) ProcessOrders(OrderManagementService_ProcessOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}
func (UnimplementedOrderManagementServiceServer) mustEmbedUnimplementedOrderManagementServiceServer() {
}

// UnsafeOrderManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServiceServer will
// result in compilation errors.
type UnsafeOrderManagementServiceServer interface {
	mustEmbedUnimplementedOrderManagementServiceServer()
}

func RegisterOrderManagementServiceServer(s grpc.ServiceRegistrar, srv OrderManagementServiceServer) {
	s.RegisterService(&OrderManagementService_ServiceDesc, srv)
}

func _OrderManagementService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderingsystem.OrderManagementService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).GetOrder(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagementService_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServiceServer).SearchOrders(m, &orderManagementServiceSearchOrdersServer{stream})
}

type OrderManagementService_SearchOrdersServer interface {
	Send(*OrderResponse) error
	grpc.ServerStream
}

type orderManagementServiceSearchOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementServiceSearchOrdersServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagementService_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServiceServer).UpdateOrders(&orderManagementServiceUpdateOrdersServer{stream})
}

type OrderManagementService_UpdateOrdersServer interface {
	SendAndClose(*OrdersResponse) error
	Recv() (*OrdersRequest, error)
	grpc.ServerStream
}

type orderManagementServiceUpdateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementServiceUpdateOrdersServer) SendAndClose(m *OrdersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementServiceUpdateOrdersServer) Recv() (*OrdersRequest, error) {
	m := new(OrdersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagementService_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServiceServer).ProcessOrders(&orderManagementServiceProcessOrdersServer{stream})
}

type OrderManagementService_ProcessOrdersServer interface {
	Send(*OrderResponse) error
	Recv() (*OrdersRequest, error)
	grpc.ServerStream
}

type orderManagementServiceProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementServiceProcessOrdersServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementServiceProcessOrdersServer) Recv() (*OrdersRequest, error) {
	m := new(OrdersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementService_ServiceDesc is the grpc.ServiceDesc for OrderManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderingsystem.OrderManagementService",
	HandlerType: (*OrderManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderManagementService_GetOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchOrders",
			Handler:       _OrderManagementService_SearchOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateOrders",
			Handler:       _OrderManagementService_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ProcessOrders",
			Handler:       _OrderManagementService_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/proto/orderingsystem.proto",
}
