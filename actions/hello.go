package actions

import "github.com/garrus/go-blog/core"

func Hello(c *core.Context){

	c.SetHeader("x-blogo-version", "v0.1")
	c.Send("<h1>Hello world!</h1>", 200)
}