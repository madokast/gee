package a01_std_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/madokast/gee/utils"
)

func A01_std_http() {
	// 监听 /hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		size, _ := w.Write([]byte("A01_std_http " + time.Now().Format("2006-01-02 15:04:05")))
		fmt.Println("Call from ", r.URL, " write ", size, " B")
	})
	// 采用协程，这样 main 退出调用了 exit 所有协程都会终止
	go http.ListenAndServe(":8080", nil)

	fmt.Println("Start server")

	// 循环调用方法，Get 请求
	utils.TimedLoop(time.Second, time.Millisecond, 5, func() {
		fmt.Println(utils.HttpGet("http://localhost:8080/hello"))
	})
}

// Call from  /  write  32  B
// A01_std_http 2022-12-15 21:10:19
// Call from  /  write  32  B
// A01_std_http 2022-12-15 21:10:20
