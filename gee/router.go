package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]handlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]handlerFunc)}
}

func (r *router) addRouter(method string, pattern string, f handlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = f
}

func (r *router) ServeHTTP(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}