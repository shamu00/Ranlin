package ran

type router struct {
	handlers map[string]HandlerFunction
}

func newRouter() *router{
	return &router{
		handlers: make(map[string]HandlerFunction),
	}
}

func (r *router)addRoute(method, pattern string, handler HandlerFunction) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router)handle(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		// TODO: err
	}
}
