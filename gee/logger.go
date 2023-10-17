package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.Printf("[%d], %s in %v", c.StatusCode, c.Req.URL.Path, time.Since(t))
	}
}
