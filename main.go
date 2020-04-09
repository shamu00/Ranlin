package main

import (
	"fmt"
	"net/http"
	"os"
)

type rootHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Engine struct {}

func main() {
	//http.Handle("/", new(Engine))
	//http.HandleFunc("/hello", helloHandler)
	engine := new(Engine)
	fmt.Print("Server start...\n")
	err := http.ListenAndServe("localhost:23333", engine)
	if err != nil {
		panic(err)
	}
}

func (e *Engine)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Host)
	for k, v := range r.Header {
		_, err := fmt.Fprintf(w, "%q=%q\n", k, v)
		if err != nil {
			fmt.Printf("err:%v", err)
			os.Exit(-1)
		}
	}

}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	for k, v := range r.Header {
		_, err := fmt.Fprintf(w,"%q:%q\n", k, v)
		if err != nil {
			fmt.Printf("err:%v", err)
			os.Exit(-1)
		}
	}

}