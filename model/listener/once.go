package listener

import (
	"go-engine.io/model"
)

type Once struct {
	callback model.Callback
}

func NewOnceListener(callback model.Callback) Once {
	return Once{callback: callback}
}

func (o Once) Call(args ...interface{}) model.Callback {
	o.callback(args)

	return o.callback
}
