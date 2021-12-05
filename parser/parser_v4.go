package parser

import (
	"go-engine.io/model"
)

type V4 struct {}

func NewParserV4() V4 {
	panic("implement me")
}

func (p V4) DecodePacket(data interface{}) Packet {
	panic("implement me")
}

func (p V4) EncodePacket(packet Packet, callback model.Callback) {
	panic("implement me")
}

func (p V4) EncodePayLoad(packets []Packet, callbackFunc model.Callback) {
	panic("implement me")
}

func (p V4) DecodePayload(data interface{}, callbackFunc model.Callback) {
	panic("implement me")
}

func (p V4) GetProtocolVersion() int {
	return 4
}
