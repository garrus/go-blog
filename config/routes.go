package config

import (
	"github.com/garrus/go-blog/core"
	"github.com/garrus/go-blog/actions"
)

func SetRoutes(r *core.Router) {

	r.AddRoute("GET", "/hello", actions.Hello)
}