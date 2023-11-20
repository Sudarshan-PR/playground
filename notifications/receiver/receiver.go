package receiver

import (
	"context"
	"errors"
	"log"

	pb "github.com/Sudarshan-PR/playground/notifications/notification-protos"
	"github.com/Sudarshan-PR/playground/notifications/sender"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServer
}

func (r *NotificationServer) PushNotification (ctx context.Context,req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	//var data map[string]interface{}
	userid := req.Userid
	if userid == "" {
		return nil, errors.New("Invalid userid sent")
	}
	
	log.Printf("Message Received: ", req)
	//dataMarshalled, _ := json.Marshal(req)
	//_ = json.Unmarshal(dataMarshalled, &data)
	
	//_ = sender.SendToUser(userid, data)
	if err := sender.SendToUser(userid, req); err != nil {
		return nil, errors.New("Error while sending message to user")
	}

	return &pb.NotificationResponse{Sent: true, Error: ""}, nil
}
