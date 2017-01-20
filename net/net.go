package net

import "github.com/gorilla/websocket"

type PacketID uint8

type OutboundPacket interface {
	Send(conn *websocket.Conn)
}

type InboundPacket interface {

}