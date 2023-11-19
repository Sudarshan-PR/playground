package receiver

import (
	"context"

	pb "github.com/Sudarshan-PR/playground/notifications/protos"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServer
}

func (r *NotificationServer) PushNotification (context.Context, *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	return &pb.NotificationResponse{}, nil
}
