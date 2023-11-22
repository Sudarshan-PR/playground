package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Sudarshan-PR/playground/notifications/notification-protos"
	"github.com/Sudarshan-PR/playground/notifications/receiver"
	"github.com/Sudarshan-PR/playground/notifications/sender"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

// Setting up websocket server
var upgrader = websocket.Upgrader{} // use default options

func serveGRPC(server *grpc.Server, listener net.Listener) {
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
func websocketHandler(w http.ResponseWriter, req *http.Request) {
	userid := req.URL.Query().Get("user")
	if userid == "" {
		log.Printf("No userid passed.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No user passed as query param"))
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket:", err)
		return
	}
	defer c.Close()
	
	// Register userid in system
	err = sender.RegisterClient(userid, c)
	if err != nil {
		log.Printf("Error registering to websocket:", err)
		return
	}
	defer sender.UnregisterClient(userid)
	

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Disconnected: ", userid,  err)
			break
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func serveWebSocket() {
	http.HandleFunc("/ws", websocketHandler)
	fmt.Println("Starting WebSocket Server at 0.0.0.0:8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Failed to start websocket server: %v", err)
	}
}

func main() {
	// Setting up gRPC server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	serverGRPC := grpc.NewServer()
	pb.RegisterNotificationServer(serverGRPC, &receiver.NotificationServer{})
	log.Printf("gRPC listening at %v", lis.Addr())
	
	forever := make(chan bool)
	
	go serveGRPC(serverGRPC, lis)
	go serveWebSocket()

	<- forever
}
