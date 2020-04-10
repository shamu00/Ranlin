package main

import (
	"fmt"
	"net/http"
	"ran"
)

func main() {
	ranlin := ran.New()
	ranlin.GET("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "%q=%q", k, v)
		}
	})
	ranlin.Run("localhost:2333")
}