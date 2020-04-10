package ran

import (
	"fmt"
	"net/http"
)

type HandlerFunction func(http.ResponseWriter, *http.Request)

const (
	M_POST = "POST"
	M_GET = "GET"
)
type ranEngine struct {
	router map[string]HandlerFunction
}

func New() *ranEngine {
	engine := new(ranEngine)
	engine.router = make(map[string]HandlerFunction)
	return engine
}

func (ran *ranEngine) addRoute(method string, pattern string, handler HandlerFunction) {
	str := method + pattern
	//TODO: log
	ran.router[str] = handler
}

func (ran *ranEngine) GET(pattern string, handler HandlerFunction) {
	ran.addRoute(M_GET, pattern, handler)
}

func (ran *ranEngine) POST(pattern string, handler HandlerFunction) {
	ran.addRoute(M_POST, pattern, handler)
}

func (ran *ranEngine) Run(addr string) error{
	return http.ListenAndServe(addr, ran)
}

func (ran *ranEngine)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + r.URL.Path
	if handler, ok := ran.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND:%s", r.URL)
	}
}