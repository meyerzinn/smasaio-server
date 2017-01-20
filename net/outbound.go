package net

import (
	"github.com/gogo/protobuf/test/custom"
	"github.com/gorilla/websocket"
)

const (
	OUTBOUND_WORLD_SNAPSHOT PacketID = iota + 128
)

type OutboundWorldSnapshot struct {

}

func (OutboundWorldSnapshot) Send(conn *websocket.Conn) {
	panic("implement me")
}

