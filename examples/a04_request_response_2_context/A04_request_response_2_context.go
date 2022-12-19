package a04_request_response_2_context

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/madokast/gee/utils"
)

type Context struct {
	// 原始的请求和响应
	ResWriter http.ResponseWriter
	Request   *http.Request
	// 提取到 Context 中
	Path   string
	Method string
	// 返回状态
	StatusCode int
}
type JSONOBJ = map[string]interface{}
type Handler = func(*Context)
type router struct{ handlerMap map[string]Handler } // 路由暂时用 map 实现
type Engine struct{ r *router }

/*========================== Context =========================================*/
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{ResWriter: w, Request: r, Path: r.URL.Path, Method: r.Method}
}
func (c *Context) PostForm(key string) string  { return c.Request.FormValue(key) }       // 获取表单值
func (c *Context) Query(key string) string     { return c.Request.URL.Query().Get(key) } // Get 请求的 Query
func (c *Context) SetHeader(key, value string) { c.ResWriter.Header().Set(key, value) }
func (c *Context) Status(code int)             { c.ResWriter.WriteHeader(code); c.StatusCode = code }
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("content-type", "application/json")
	c.Status(code)
	json.NewEncoder(c.ResWriter).Encode(obj)
}
func (c *Context) HTML(code int, html string) {
	c.SetHeader("content-type", "text/html")
	c.Status(code)
	c.ResWriter.Write([]byte(html))
}

/*========================== router =========================================*/
func NewRouter() *router { return &router{make(map[string]func(*Context))} }
func (r *router) addRoute(method, pattern string, handler Handler) {
	r.handlerMap[method+"-"+pattern] = handler
}
func (r *router) handle(c *Context) {
	if handler, ok := r.handlerMap[c.Method+"-"+c.Path]; ok {
		handler(c)
	} else {
		c.HTML(http.StatusNotFound, "<h1>Not Found</h1>")
	}
}

/*========================== engine =========================================*/
func NewEngine() *Engine                                           { return &Engine{r: NewRouter()} }
func (e *Engine) addRoute(method, pattern string, h Handler)       { e.r.addRoute(method, pattern, h) }
func (e *Engine) GET(pattern string, handler Handler)              { e.addRoute("GET", pattern, handler) }
func (e *Engine) POST(pattern string, handler Handler)             { e.addRoute("POST", pattern, handler) }
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) { e.r.handle(NewContext(w, r)) } // 实现 http.Handler
func (e *Engine) Run(addr string)                                  { http.ListenAndServe(addr, e) }

/*========================== 测试 =========================================*/
func Test() {
	engine := NewEngine()
	engine.GET("/get", func(c *Context) { c.JSON(http.StatusOK, JSONOBJ{"name": c.Query("name")}) })
	engine.POST("/post", func(c *Context) { c.HTML(http.StatusOK, "<h1>Age="+c.PostForm("age")+"</h1>") })

	go engine.Run(":8080")

	// {"name":"mdk"}
	// <h1>14</h1>
	utils.TimedLoop(time.Second, time.Millisecond, 5, func() {
		// 发送 query
		println(utils.HttpGet("http://localhost:8080/get?name=mdk&age=14"))
		// 发送表单
		println(utils.HttpPOST("http://localhost:8080/post", "application/x-www-form-urlencoded", "name=mdk&age=14"))
	})
}
