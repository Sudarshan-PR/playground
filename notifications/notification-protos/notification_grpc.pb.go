// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notification.proto

package notifications_protos

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

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	PushNotification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) PushNotification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/Notification/PushNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
// All implementations must embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	PushNotification(context.Context, *NotificationRequest) (*NotificationResponse, error)
	mustEmbedUnimplementedNotificationServer()
}

// UnimplementedNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) PushNotification(context.Context, *NotificationRequest) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNotification not implemented")
}
func (UnimplementedNotificationServer) mustEmbedUnimplementedNotificationServer() {}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_PushNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).PushNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notification/PushNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).PushNotification(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushNotification",
			Handler:    _Notification_PushNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}