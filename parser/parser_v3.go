package parser

import (
	"go-engine.io/model"
)

type V3 struct{}

func NewParserV3() V3 {
	panic("implement me")
}

func (p V3) DecodePacket(data interface{}) Packet {
	panic("implement me")
}

func (p V3) EncodePacket(packet Packet, callback model.Callback) {
	panic("implement me")
}

func (p V3) EncodePayLoad(packets []Packet, callbackFunc model.Callback) {
	panic("implement me")
}

func (p V3) DecodePayload(data interface{}, callbackFunc model.Callback) {
	panic("implement me")
}

func (p V3) GetProtocolVersion() int {
	return 3
}
