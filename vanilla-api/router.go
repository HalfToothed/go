package main

import (
	"database/sql"
	"net/http"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(db *sql.DB) *router {
	router := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	service := NewService(db)

	router.addRoute("POST", "/products", service.CreateProduct)
	router.addRoute("GET", "/products", service.GetProduct)

	return router
}

func (r *router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}
