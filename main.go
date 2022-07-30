package main

import (
	"gee-web/gee"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.String(200, "hello world")
	})

	r.Get("/duck", func(c *gee.Context) {
		c.String(200, "duck")
	})

	r.Run("localhost:8080")
}
