package listener

import "go-engine.io/model"

type Default struct {
	callback model.Callback
}

func NewDefaultListener(callback model.Callback) Default {
	return Default{callback: callback}
}

func (d Default) Call(args ...interface{}) model.Callback {
	d.callback(args)

	return d.callback
}
