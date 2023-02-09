package fetch

import (
	"bytes"
	_ "embed"
	"github.com/dop251/goja"
	"github.com/monkeyWie/gopeed/pkg/download/extension/polyfill"
	"io"
	"net/http"
)

//go:embed fetch.js
var script string

type Response struct {
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
}

func Enable(vm *goja.Runtime) error {
	prg, err := goja.Compile("fetch.js", script, false)
	if err != nil {
		return err
	}
	vm.Set("__fetch__", getFetch(vm))
	vm.RunProgram(prg)
	return nil
}

func getFetch(vm *goja.Runtime) any {
	// fetch inner function
	return func(method string, url string, body string, headers map[string]string) *Response {
		if method == "" {
			method = "GET"
		}
		var b io.Reader
		if body != "" {
			b = bytes.NewBufferString(body)
		}
		req, err := http.NewRequest(method, url, b)
		if err != nil {
			polyfill.Throw(vm, err.Error())
		}
		if len(headers) > 0 {
			for k, v := range headers {
				req.Header.Set(k, v)
			}
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			polyfill.Throw(vm, err.Error())
		}
		defer resp.Body.Close()
		var buf bytes.Buffer
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			polyfill.Throw(vm, err.Error())
		}
		respWrap := &Response{
			Status:     resp.StatusCode,
			StatusText: resp.Status,
			Body:       buf.String(),
		}
		if len(resp.Header) > 0 {
			respWrap.Headers = make(map[string]string)
			for k := range resp.Header {
				respWrap.Headers[k] = resp.Header.Get(k)
			}
		}
		return respWrap
	}
}
