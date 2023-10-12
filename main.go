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

	v1 := r.Group("/v1")
	{
		v1.Get("/hello", func(c *gee.Context) {
			c.String(200, "hello %s,you are at %s", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("v2")
	{
		v2.Get("/thank", func(c *gee.Context) {
			c.String(200, "thank %s,you are at %s", c.Query("name"), c.Path)
		})
	}

	r.Get("/hello/:name", func(c *gee.Context) {
		c.String(200, "hello %s,you are at %s", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"filepath": c.Params["filepath"],
		})
	})

	r.Run(":8080")
}
