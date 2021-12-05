package server

import (
	"go-engine.io/model/emitter"
	"go-engine.io/transport"
	"net/http"
	"sync"
)

type socket struct {
	currentTransPortName string
	query                map[string]string
	headers              map[string][]string
	transport            transport.Interface
	emitter              emitter.Emitter
}

func NewEngineIOSocket(mutex *sync.Mutex, sid string, parserVersion int) socket {
	panic("implement me")
}

func (s *socket) Init(transport transport.Interface) {
	s.query = transport.GetInitialQuery()
	s.headers = transport.GetInitialHeaders()
	s.onOpen()
}

func (s *socket) UpdateInitialHeadersFromActiveTransport() {
	s.query = s.transport.GetInitialQuery()
	s.headers = s.transport.GetInitialHeaders()
}

func (s socket) onOpen() {

}

func (s socket) OnRequest(r *http.Request, w http.ResponseWriter) string {
	panic("implement me")
}
func (s socket) GetCurrentTransportName() string {
	return s.currentTransPortName
}
