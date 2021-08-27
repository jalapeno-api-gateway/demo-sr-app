// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pushservice

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

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushServiceClient interface {
	SubscribeToDataRates(ctx context.Context, in *DataRateSubscription, opts ...grpc.CallOption) (PushService_SubscribeToDataRatesClient, error)
	SubscribeToLsNodes(ctx context.Context, in *LsNodeSubscription, opts ...grpc.CallOption) (PushService_SubscribeToLsNodesClient, error)
	SubscribeToLsLinks(ctx context.Context, in *LsLinkSubscription, opts ...grpc.CallOption) (PushService_SubscribeToLsLinksClient, error)
	SubscribeToTotalPacketsSent(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (PushService_SubscribeToTotalPacketsSentClient, error)
	SubscribeToTotalPacketsReceived(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (PushService_SubscribeToTotalPacketsReceivedClient, error)
}

type pushServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushServiceClient(cc grpc.ClientConnInterface) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) SubscribeToDataRates(ctx context.Context, in *DataRateSubscription, opts ...grpc.CallOption) (PushService_SubscribeToDataRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PushService_ServiceDesc.Streams[0], "/pushservice.PushService/SubscribeToDataRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeToDataRatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeToDataRatesClient interface {
	Recv() (*DataRateEvent, error)
	grpc.ClientStream
}

type pushServiceSubscribeToDataRatesClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeToDataRatesClient) Recv() (*DataRateEvent, error) {
	m := new(DataRateEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushServiceClient) SubscribeToLsNodes(ctx context.Context, in *LsNodeSubscription, opts ...grpc.CallOption) (PushService_SubscribeToLsNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PushService_ServiceDesc.Streams[1], "/pushservice.PushService/SubscribeToLsNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeToLsNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeToLsNodesClient interface {
	Recv() (*LsNodeEvent, error)
	grpc.ClientStream
}

type pushServiceSubscribeToLsNodesClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeToLsNodesClient) Recv() (*LsNodeEvent, error) {
	m := new(LsNodeEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushServiceClient) SubscribeToLsLinks(ctx context.Context, in *LsLinkSubscription, opts ...grpc.CallOption) (PushService_SubscribeToLsLinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &PushService_ServiceDesc.Streams[2], "/pushservice.PushService/SubscribeToLsLinks", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeToLsLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeToLsLinksClient interface {
	Recv() (*LsLinkEvent, error)
	grpc.ClientStream
}

type pushServiceSubscribeToLsLinksClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeToLsLinksClient) Recv() (*LsLinkEvent, error) {
	m := new(LsLinkEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushServiceClient) SubscribeToTotalPacketsSent(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (PushService_SubscribeToTotalPacketsSentClient, error) {
	stream, err := c.cc.NewStream(ctx, &PushService_ServiceDesc.Streams[3], "/pushservice.PushService/SubscribeToTotalPacketsSent", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeToTotalPacketsSentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeToTotalPacketsSentClient interface {
	Recv() (*TelemetryEvent, error)
	grpc.ClientStream
}

type pushServiceSubscribeToTotalPacketsSentClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeToTotalPacketsSentClient) Recv() (*TelemetryEvent, error) {
	m := new(TelemetryEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushServiceClient) SubscribeToTotalPacketsReceived(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (PushService_SubscribeToTotalPacketsReceivedClient, error) {
	stream, err := c.cc.NewStream(ctx, &PushService_ServiceDesc.Streams[4], "/pushservice.PushService/SubscribeToTotalPacketsReceived", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeToTotalPacketsReceivedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeToTotalPacketsReceivedClient interface {
	Recv() (*TelemetryEvent, error)
	grpc.ClientStream
}

type pushServiceSubscribeToTotalPacketsReceivedClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeToTotalPacketsReceivedClient) Recv() (*TelemetryEvent, error) {
	m := new(TelemetryEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushServiceServer is the server API for PushService service.
// All implementations must embed UnimplementedPushServiceServer
// for forward compatibility
type PushServiceServer interface {
	SubscribeToDataRates(*DataRateSubscription, PushService_SubscribeToDataRatesServer) error
	SubscribeToLsNodes(*LsNodeSubscription, PushService_SubscribeToLsNodesServer) error
	SubscribeToLsLinks(*LsLinkSubscription, PushService_SubscribeToLsLinksServer) error
	SubscribeToTotalPacketsSent(*TelemetrySubscription, PushService_SubscribeToTotalPacketsSentServer) error
	SubscribeToTotalPacketsReceived(*TelemetrySubscription, PushService_SubscribeToTotalPacketsReceivedServer) error
	mustEmbedUnimplementedPushServiceServer()
}

// UnimplementedPushServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushServiceServer struct {
}

func (UnimplementedPushServiceServer) SubscribeToDataRates(*DataRateSubscription, PushService_SubscribeToDataRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToDataRates not implemented")
}
func (UnimplementedPushServiceServer) SubscribeToLsNodes(*LsNodeSubscription, PushService_SubscribeToLsNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToLsNodes not implemented")
}
func (UnimplementedPushServiceServer) SubscribeToLsLinks(*LsLinkSubscription, PushService_SubscribeToLsLinksServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToLsLinks not implemented")
}
func (UnimplementedPushServiceServer) SubscribeToTotalPacketsSent(*TelemetrySubscription, PushService_SubscribeToTotalPacketsSentServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToTotalPacketsSent not implemented")
}
func (UnimplementedPushServiceServer) SubscribeToTotalPacketsReceived(*TelemetrySubscription, PushService_SubscribeToTotalPacketsReceivedServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToTotalPacketsReceived not implemented")
}
func (UnimplementedPushServiceServer) mustEmbedUnimplementedPushServiceServer() {}

// UnsafePushServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServiceServer will
// result in compilation errors.
type UnsafePushServiceServer interface {
	mustEmbedUnimplementedPushServiceServer()
}

func RegisterPushServiceServer(s grpc.ServiceRegistrar, srv PushServiceServer) {
	s.RegisterService(&PushService_ServiceDesc, srv)
}

func _PushService_SubscribeToDataRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataRateSubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).SubscribeToDataRates(m, &pushServiceSubscribeToDataRatesServer{stream})
}

type PushService_SubscribeToDataRatesServer interface {
	Send(*DataRateEvent) error
	grpc.ServerStream
}

type pushServiceSubscribeToDataRatesServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeToDataRatesServer) Send(m *DataRateEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _PushService_SubscribeToLsNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LsNodeSubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).SubscribeToLsNodes(m, &pushServiceSubscribeToLsNodesServer{stream})
}

type PushService_SubscribeToLsNodesServer interface {
	Send(*LsNodeEvent) error
	grpc.ServerStream
}

type pushServiceSubscribeToLsNodesServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeToLsNodesServer) Send(m *LsNodeEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _PushService_SubscribeToLsLinks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LsLinkSubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).SubscribeToLsLinks(m, &pushServiceSubscribeToLsLinksServer{stream})
}

type PushService_SubscribeToLsLinksServer interface {
	Send(*LsLinkEvent) error
	grpc.ServerStream
}

type pushServiceSubscribeToLsLinksServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeToLsLinksServer) Send(m *LsLinkEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _PushService_SubscribeToTotalPacketsSent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TelemetrySubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).SubscribeToTotalPacketsSent(m, &pushServiceSubscribeToTotalPacketsSentServer{stream})
}

type PushService_SubscribeToTotalPacketsSentServer interface {
	Send(*TelemetryEvent) error
	grpc.ServerStream
}

type pushServiceSubscribeToTotalPacketsSentServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeToTotalPacketsSentServer) Send(m *TelemetryEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _PushService_SubscribeToTotalPacketsReceived_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TelemetrySubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).SubscribeToTotalPacketsReceived(m, &pushServiceSubscribeToTotalPacketsReceivedServer{stream})
}

type PushService_SubscribeToTotalPacketsReceivedServer interface {
	Send(*TelemetryEvent) error
	grpc.ServerStream
}

type pushServiceSubscribeToTotalPacketsReceivedServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeToTotalPacketsReceivedServer) Send(m *TelemetryEvent) error {
	return x.ServerStream.SendMsg(m)
}

// PushService_ServiceDesc is the grpc.ServiceDesc for PushService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pushservice.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToDataRates",
			Handler:       _PushService_SubscribeToDataRates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToLsNodes",
			Handler:       _PushService_SubscribeToLsNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToLsLinks",
			Handler:       _PushService_SubscribeToLsLinks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToTotalPacketsSent",
			Handler:       _PushService_SubscribeToTotalPacketsSent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToTotalPacketsReceived",
			Handler:       _PushService_SubscribeToTotalPacketsReceived_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pushservice.proto",
}