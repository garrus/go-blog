package core

import (
	"strings"
)

type ActionFunc func(*Context)

type Route struct {
	method string
	prefix string
	handler ActionFunc
}


type Router struct {
	routes []Route
}


func (r *Router) AddRoute(method string, prefix string, handler ActionFunc) {
	r.routes = append(r.routes, Route{strings.ToUpper(method), strings.ToLower(prefix), handler})
}


func (r *Router) Route(method string, uri string) (ActionFunc, bool) {

	uri = strings.ToLower(uri)

	for _, route := range r.routes {
		if route.method != "" && route.method != method {
			continue
		}
		if !strings.HasPrefix(uri, route.prefix) {
			continue
		}
		return route.handler, true
	}
	return nil, false
}