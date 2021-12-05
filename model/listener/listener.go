package listener

import "go-engine.io/model"

type Interface interface {
	Call(args... interface{}) model.Callback
}
