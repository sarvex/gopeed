package extension

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/url"
	"github.com/monkeyWie/gopeed/pkg/download/extension/polyfill/fetch"
	"os"
)

func main() {
	file, err := os.ReadFile("D:\\code\\study\\node\\gopeed-extension-test\\dist\\index.js")
	if err != nil {
		panic(err)
	}
	SCRIPT := string(file)

	loop := eventloop.NewEventLoop()
	loop.Run(func(vm *goja.Runtime) {
		url.Enable(vm)
		if err := fetch.Enable(vm); err != nil {
			panic(err)
		}
		SCRIPT = `
(async function(){
	const resp = await fetch('https://www.baidu.com')
	console.log(resp.text())
})()
`
		_, err2 := vm.RunString(SCRIPT)
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		//vm.RunProgram(prg)
	})
}
