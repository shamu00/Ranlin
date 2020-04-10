package ran

import (
	"net/http"
)

type HandlerFunction func(c *Context)

const (
	MPost = "POST"
	MGet = "GET"
)

type ranEngine struct {
	router *router
}

func New() *ranEngine {
	return &ranEngine{newRouter()}
}

func (ran *ranEngine) addRoute(method, pattern string, handler HandlerFunction) {
	ran.router.addRoute(method, pattern, handler)
}
func (ran *ranEngine) GET(pattern string, handler HandlerFunction) {
	ran.addRoute(MGet, pattern, handler)
}

func (ran *ranEngine) POST(pattern string, handler HandlerFunction) {
	ran.addRoute(MPost, pattern, handler)
}

func (ran *ranEngine) Run(addr string) error{
	return http.ListenAndServe(addr, ran)
}

func (ran *ranEngine)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	ran.router.handle(c)
}