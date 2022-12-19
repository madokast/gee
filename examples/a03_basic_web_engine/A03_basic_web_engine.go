package a03_basic_web_engine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/madokast/gee/utils"
)

const (
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
)

type HanderFunc = func(http.ResponseWriter, *http.Request)

type Engine struct {
	// 路由
	router map[string]HanderFunc
}

func NewEnginc() *Engine {
	return &Engine{router: make(map[string]HanderFunc)}
}

// 添加路由
func (enginc *Engine) addRoute(method string, pattern string, handler HanderFunc) {
	key := method + "-" + pattern
	enginc.router[key] = handler
}

func (enginc *Engine) GET(pattern string, handler HanderFunc) {
	enginc.addRoute(METHOD_GET, pattern, handler)
}

func (enginc *Engine) POST(pattern string, handler HanderFunc) {
	enginc.addRoute(METHOD_POST, pattern, handler)
}

func (enginc *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := enginc.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", key)
	}
}

func (enginc *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, enginc)
}

/*===================== 测试 ==========================*/

func A03_basic_web_engine() {
	enginc := NewEnginc()
	enginc.GET("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "GET /") })
	enginc.POST("/hello", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "POST /hello") })

	go enginc.Run(":8080")

	// GET /
	// POST /hello
	// 404 NOT FOUND: GET-/world
	utils.TimedLoop(time.Second, time.Millisecond, 5, func() {
		fmt.Println(utils.HttpGet("http://localhost:8080/"))
		fmt.Println(utils.HttpPOST("http://localhost:8080/hello", "", ""))
		fmt.Println(utils.HttpGet("http://localhost:8080/world"))
	})
}
