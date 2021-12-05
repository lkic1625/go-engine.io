package emitter

import (
	"go-engine.io/model"
	"go-engine.io/model/listener"
)

type Interface interface {
	Emit(eventName string, args... interface{}) Interface
	On(eventName string, callback model.Callback) Interface
	Off(eventName string) Interface
	Once(eventName string, callback model.Callback) Interface
}

type Emitter struct {
	callbacks map[string]model.CallbackLinkedQueue
}

func (e Emitter) Emit(eventName string, args... interface{}) Interface {
	for _, callback := range e.callbacks[eventName] {
		callback.Call(args)
		if _, ok := callback.(listener.Once); ok {
			e.Off(eventName)
		}
	}

	return e
}

func (e Emitter) Off(eventName string) Interface {
	e.callbacks[eventName] = []listener.Interface{}

	return e
}

func (e Emitter) On(eventName string, callback model.Callback) Interface {
	return e.on(eventName, listener.NewDefaultListener(callback))
}

func (e Emitter) Once(eventName string, callback model.Callback) Interface {
	return e.on(eventName, listener.NewOnceListener(callback))
}

func (e Emitter) on(eventName string, li listener.Interface) Interface {
	if callbacks, ok := e.callbacks[eventName]; ok {
		e.callbacks[eventName] = []listener.Interface{}
	} else {
		callbacks = append(callbacks, li)
	}

	return e
}
