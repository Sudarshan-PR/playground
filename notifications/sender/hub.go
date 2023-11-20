package sender

import (
	"errors"

	"github.com/gorilla/websocket"
)

// Holds all websocket connections
var connectionHub = make(map[string]*websocket.Conn)

func RegisterClient(userid string, conn *websocket.Conn) error {
	_, exists := connectionHub[userid]
	if exists {
		return errors.New("User ID already exists")
	}

	connectionHub[userid] = conn
	return nil
}

func UnregisterClient(userid string) {
	delete(connectionHub, userid)
}

func SendToUser(userid string, data interface{}) error {
	conn, ok := connectionHub[userid]
	if !ok {
		return errors.New("No connection for user found")
	}
	err := conn.WriteJSON(data)
	if err != nil {
		return errors.New("Error while writing to connection")
	}

	return nil
}
