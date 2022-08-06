package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.String(http.StatusOK, "gugugu")
		})

		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "gu!")
		})
	}

	r.Run(":8080")
}
