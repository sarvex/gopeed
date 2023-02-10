package extension

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/url"
	"github.com/monkeyWie/gopeed/pkg/download/extension/polyfill/fetch"
)

type Engine struct {
	loop *eventloop.EventLoop
}

func (e *Engine) Run(script string) (err error) {
	e.loop.Run(func(vm *goja.Runtime) {
		url.Enable(vm)
		if err = fetch.Enable(vm); err != nil {
			return
		}
		_, err = vm.RunString(script)
		return
	})
	return
}

func NewEngine() *Engine {
	loop := eventloop.NewEventLoop()
	return &Engine{
		loop: loop,
	}
}

func Run(script string) (err error) {
	engine := NewEngine()
	return engine.Run(script)
}
