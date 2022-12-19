package main

import (
	a01 "github.com/madokast/gee/examples/a01_std_http"
	a02 "github.com/madokast/gee/examples/a02_std_http_handler"
	a03 "github.com/madokast/gee/examples/a03_basic_web_engine"
	a04 "github.com/madokast/gee/examples/a04_request_response_2_context"
)

func main() {
	examples := []func(){
		a01.A01_std_http,
		a02.A02_std_http_handler,
		a03.A03_basic_web_engine,
		a04.Test,
	}

	examples[3]()
}
