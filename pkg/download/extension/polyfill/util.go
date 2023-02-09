package polyfill

import (
	"errors"
	"github.com/dop251/goja"
)

func Throw(vm *goja.Runtime, msg string) {
	panic(vm.NewGoError(errors.New(msg)))
}
