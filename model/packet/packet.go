package packet

type Interface interface {
	getDataAsString() string
	getDataAsByte() []byte
}

type BytePacket struct {
	PacketType Type
	Data       []byte
}

type StringPacket struct {
	PacketType Type
	Data       string
}
