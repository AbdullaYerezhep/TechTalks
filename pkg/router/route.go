package router

import (
	"errors"
	"net/http"
	"regexp"
)

type route struct {
	Path    string
	Method  string
	Handler Handler
}

type Router struct {
	routes []route
}
type Handler func(http.ResponseWriter, *http.Request)

// add method, path, handler func in to slice of route in struct Router
func (r *Router) AddRoute(method, path string, handler Handler) {
	r.routes = append(r.routes, route{Path: path, Method: method, Handler: handler})
}

// methods: Get, Post, Put, Delete of router created for use in other package
func (r *Router) Get(path string, handler Handler) {
	r.AddRoute(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler Handler) {
	r.AddRoute(http.MethodPost, path, handler)
}

func (r *Router) Put(path string, handler Handler) {
	r.AddRoute(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handler Handler) {
	r.AddRoute(http.MethodDelete, path, handler)
}

// from request of client we find handler function if it is exist in Router
func (r *Router) getHandler(path, method string) (Handler, error) {
	for _, route := range r.routes {

		regex := regexp.MustCompile(route.Path)

		if regex.MatchString(path) && route.Method == method {
			return route.Handler, nil
		}
	}
	return nil, errors.New("not foundm")
}

// for implementing http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	handler, err := r.getHandler(path, method)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	handler(w, req)
}

// if call this function you get type *Router, and use all method of Router type
func NewRouter() *Router {
	return &Router{}
}
