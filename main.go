package main

import (
	"Gee/gee"
)

func main() {
	r := gee.New()
	//r.Get("/", func(c *gee.Context) {
	//	c.HTML(200, "<h1>hello</h1>")
	//})
	//
	//r.Get("/hello", func(c *gee.Context) {
	//	c.String(200, "hello %s,you are at %s", c.Query("name"), c.Path)
	//})
	//
	//r.POST("/login", func(c *gee.Context) {
	//	c.JSON(200, gee.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})

	r.Get("/hello/:name", func(c *gee.Context) {
		c.String(200, "hello %s,you are at %s", c.Param("name"), c.Path)
	})

	r.Run(":8080")
}
