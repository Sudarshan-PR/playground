package main

import (
	"log"
	"net"

	pb "github.com/Sudarshan-PR/playground/notifications/protos"
	"github.com/Sudarshan-PR/playground/notifications/receiver"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterNotificationServer(server, &receiver.NotificationServer{})
	log.Printf("Listening at port: %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
