package ran

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin
	Request *http.Request
	Writer http.ResponseWriter
	// req info
	Path string
	Method string
	// rsp info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context {
		Request: r,
		Writer: w,
		Path: r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context)HTML(code int, html string, values ...interface{}) {
	c.setHeader("Content-Type", "text/html")
	c.Status(http.StatusOK)
	c.Writer.Write([]byte(html))
}

func (c *Context)String(code int, formatStr string, values ...interface{}) {
	c.setHeader("Content-Type", "text/plain")
	c.Status(http.StatusOK)
	c.Writer.Write([]byte(fmt.Sprintf(formatStr, values...)))
}

func (c *Context)JSON(code int, obj interface{}) {
	c.setHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context)Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context)Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context)Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context)setHeader(key, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context)PostForm(key string) string{
	return c.Request.FormValue(key)
}