package transport

import "net/http"

type Interface interface {
	GetInitialQuery() map[string]string
	GetInitialHeaders() map[string][]string
	OnRequest(r *http.Request, w http.ResponseWriter)
	OnData()
	OnClose()
	onPollRequest()
	onDataRequest()
	Send()
}