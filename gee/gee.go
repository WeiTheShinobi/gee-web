package gee

import (
	"net/http"
)

type handlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method string, url string, f handlerFunc) {
	e.router.addRouter(method, url, f)
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
	c := newContext(w, r)
	e.router.ServeHTTP(c)
}
