package main

import (
	"net/http"
	"ran"
)

func main() {
	r := ran.New()
	r.GET("/", func(c *ran.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Ran</h1>")
	})
	r.GET("/hello", func(c *ran.Context) {
		c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ran.Context) {
		c.JSON(http.StatusOK, ran.H{
			"username": c.PostForm("username"),
			"passwd": c.PostForm("passwd"),
		})
	})

	r.Run(":2333")
}