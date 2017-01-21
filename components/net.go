package components

import (
	"github.com/gorilla/websocket"
	"github.com/wooga/go-entitas"
)

type NetworkComponent struct {
	Conn *websocket.Conn
}

func (NetworkComponent) Type() entitas.ComponentType {
	return Network
}
