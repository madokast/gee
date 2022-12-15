package examples

import (
	"fmt"
	"net/http"
	"time"

	"github.com/madokast/gee/utils"
)

type Engine struct{}

// 实现 Handler 接口
func (*Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("deal /"))
	case "/hello":
		w.Write([]byte("deal /hello"))
	default:
		w.Write([]byte("deal " + r.URL.Path))
	}
}

func A02_std_http_handler() {
	// 采用协程，这样 main 退出调用了 exit 所有协程都会终止
	go http.ListenAndServe(":8080", &Engine{})

	fmt.Println("Start server")

	// 循环调用方法，Get 请求
	utils.TimedLoop(time.Second, time.Second, 5, func() {
		fmt.Println(utils.HttpGet("http://localhost:8080/"))
		fmt.Println(utils.HttpGet("http://localhost:8080/hello"))
		fmt.Println(utils.HttpGet("http://localhost:8080/world"))
	})
}

// deal /
// deal /hello
// deal /world
