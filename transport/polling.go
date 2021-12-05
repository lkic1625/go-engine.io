package transport

import (
	"go-engine.io/parser"
	"net/http"
	"sync"
)

type Polling struct{}

func (p Polling) GetInitialQuery() map[string]string {
	panic("implement me")
}

func (p Polling) GetInitialHeaders() map[string][]string {
	panic("implement me")
}

func (p Polling) OnRequest(r *http.Request, w http.ResponseWriter) {
	panic("implement me")
}

func (p Polling) OnData() {
	panic("implement me")
}

func (p Polling) OnClose() {
	panic("implement me")
}

func (p Polling) onPollRequest() {
	panic("implement me")
}

func (p Polling) onDataRequest() {
	panic("implement me")
}

func (p Polling) Send() {
	panic("implement me")
}

func NewPolling(mutex *sync.Mutex, parser parser.Interface) Polling {
	return Polling{}
}
