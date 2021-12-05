package parser

import (
	"go-engine.io/model"
)

type Interface interface {
	EncodePacket(packet Packet, callback model.Callback)
	DecodePacket(data interface{}) Packet
	EncodePayLoad(packets []Packet, callbackFunc model.Callback)
	DecodePayload(data interface{}, callbackFunc model.Callback)
	GetProtocolVersion() int
}