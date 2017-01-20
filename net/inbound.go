package net

const (
	INBOUND_SET_NAME PacketID = iota
)

var inboundPacketRegistry = map[PacketID]InboundPacket{

}
