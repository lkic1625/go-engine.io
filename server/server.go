package server

import (
	"go-engine.io/parser"
	"go-engine.io/server/handshake_interceptor"
	"go-engine.io/transport"
	"go-engine.io/util"
	"net/http"
	"sync"
)

type Server struct {
	allowedCORSOrigins []string
	clients            map[string]socket
	idGenerator        util.IDGenerator
	options            options
	transport          transport.Interface
	handshakeInterceptor handshake_interceptor.Interface
}

var (
	ParserV3 = parser.NewParserV3()
	ParserV4 = parser.NewParserV4()
)

func (e Server) handleRequest(r *http.Request, w http.ResponseWriter) {
	var queryParams = r.URL.Query()
	// TODO(CORS): Implement CORS handling.

	t := queryParams.Get("transport")
	if t != "polling" {
		e.sendErrorMessage(w, TransPortUnknown)
	}

	var sid = queryParams.Get("sid")
	if sid != "" {
		client, err := e.clients[sid]
		if err {
			e.sendErrorMessage(w, SessionIDUnknown)
		} else if client.GetCurrentTransportName() == t {
			e.sendErrorMessage(w, BadRequest)
		} else {
			client.OnRequest(r, w)
		}
	} else {
		if r.Method != "GET" {
			e.sendErrorMessage(w, BadHandshakeMethod)
		} else if e.handshakeInterceptor != nil && !e.handshakeInterceptor.Intercept(queryParams, r.Header) {
			e.sendErrorMessage(w, BadRequest)
		} else {
			e.handshakePolling(r, w)
		}
	}
}

func (e Server) sendErrorMessage(w http.ResponseWriter, code ErrorCode) {
	http.Error(w, code.String(), int(code))
}

func (e Server) handshakePolling(r *http.Request, w http.ResponseWriter) {
	sid, _ := e.idGenerator.NewID()

	var p parser.Interface
	if r.URL.Query().Get("EIO") == "4" {
		p = ParserV4
	} else {
		p = ParserV3
	}

	mutex := &sync.Mutex{}
	s := NewEngineIOSocket(mutex, sid, p.GetProtocolVersion())
	t := transport.NewPolling(mutex, p)
	s.Init(t)
	t.OnRequest(r, w)
	s.UpdateInitialHeadersFromActiveTransport()

	e.clients[sid] = s
	//
}

//func (e Server) handshakeWebsocket(r *http.Request, w http.ResponseWriter) {
//	sid, _ := e.idGenerator.NewID()
//
//	var p parser.Interface
//	if r.URL.Query().Get("EIO") == "4" {
//		p = ParserV4
//	} else {
//		p = ParserV3
//	}
//
//	// mutex behavior 잘 모름 concurrent 환경 어떻게 지원할지 고민해보기
//	socket := NewEngineIOSocket(&sync.Mutex{}, sid, p.GetProtocolVersion())
//	//
//}
