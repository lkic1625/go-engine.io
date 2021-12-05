package server

type ErrorCode int

const (
	TransPortUnknown ErrorCode = iota
	SessionIDUnknown
	BadHandshakeMethod
	BadRequest
	Forbidden
	UnsupportedProtocolVersion
)

func (e ErrorCode) String() string {
	switch e {
	case TransPortUnknown:
		return "Transport unknown"
	case SessionIDUnknown:
		return "Session ID unknown"
	case BadHandshakeMethod:
		return "Bad handshake method"
	case BadRequest:
		return "Bad request"
	case Forbidden:
		return "Forbidden"
	case UnsupportedProtocolVersion:
		return "Unsupported protocol version"
	default:
		return "Unknown error"
	}
}