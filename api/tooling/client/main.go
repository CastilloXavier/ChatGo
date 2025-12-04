package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	if err := hack1(); err != nil {
		log.Fatal(err)
	}
}

func hack1() error {
	req := make(http.Header)
	socket, _, err := websocket.DefaultDialer.Dial("http://localhost:3000/connect", req)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	defer socket.Close()

	//-----------------------------------------------------------------------------------------

	_, msg, err := socket.ReadMessage()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	if string(msg) != "HELLO" {
		return fmt.Errorf("expected HELLO, got %s", string(msg))
	}

	//-----------------------------------------------------------------------------------------

	user := struct {
		ID   uuid.UUID
		Name string
	}{
		ID:   uuid.New(),
		Name: "Xavier Castillo",
	}

	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := socket.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	//-----------------------------------------------------------------------------------------

	_, msg, err = socket.ReadMessage()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	if string(msg) != "WELCOME Xavier Castillo" {
		return fmt.Errorf("expected WELCOME, got %s", string(msg))
	}
	return nil
}
