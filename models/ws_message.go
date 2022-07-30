package models

// Struct for messages sended by websocket
type WebsocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
