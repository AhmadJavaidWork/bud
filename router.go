package main

import (
	"strings"
)

type Handler func(ResponseWriter, *Request)

type Route struct {
	Handlers map[string]Handler
	Children map[string]*Route
}

type Router struct {
	base *Route
}

func NewRouter() *Router {
	return &Router{
		base: newRoute(),
	}
}

func newRoute() *Route {
	return &Route{
		Handlers: make(map[string]Handler),
		Children: make(map[string]*Route),
	}
}

func (r *Router) addRoute(
	path string,
	method string,
	handler Handler,
) {
	splitPath := strings.Split(path, "/")
	cur := r.base

	for _, s := range splitPath {
		tmp, ok := cur.Children["/"+s]
		if !ok {
			cur.Children["/"+s] = newRoute()
			cur = cur.Children["/"+s]
		} else {
			cur = tmp
		}
	}

	cur.Handlers[method] = handler
}

func (r *Router) getHandler(path string, method string) Handler {
	splitPath := strings.Split(path, "/")
	cur := r.base

	for _, s := range splitPath {
		tmp, ok := cur.Children["/"+s]
		if !ok {
			return nil
		}
		cur = tmp
	}

	if handler, ok := cur.Handlers[method]; ok {
		return handler
	}

	return nil
}
