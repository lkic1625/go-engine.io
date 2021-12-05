package parser

import "go-engine.io/model/packet"

type Packet interface {
	getDataAsString() string
	getDataAsByte() []byte
}

type BytePacket struct {
	PacketType packet.Type
	Data       []byte
}

type StringPacket struct {
	PacketType packet.Type
	Data       string
}
