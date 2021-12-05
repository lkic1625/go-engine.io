package packet

type Type int

const (
	OPEN Type = iota
	CLOSE
	PING
	PONG
	MESSAGE
	UPGRADE
	NOOP
)

func (id Type) String() string {
	switch id {
	case OPEN:
		return "open"
	case CLOSE:
		return "close"
	case PING:
		return "ping"
	case PONG:
		return "pong"
	case MESSAGE:
		return "message"
	case UPGRADE:
		return "upgrade"
	case NOOP:
		return "noop"
	}
	return "unknown"
}