package gee

import (
	"fmt"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]handlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]handlerFunc)}
}

func (e *Engine) addRouter(method string, url string, f handlerFunc) {
	key := method + "-" + url
	e.router[key] = f
}

func (e *Engine) Get(url string, f handlerFunc) {
	e.addRouter("GET", url, f)
}

func (e *Engine) Post(url string, f handlerFunc) {
	e.addRouter("POST", url, f)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 page not found : %s\n", r.URL.Path)
	}
}
