package main

import (
	"litttlebear/auth/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter will gen router for http server
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"UserList", "GET", "/users", handler.UsersGetHandler,
	},
	Route{
		"UserInfo", "GET", "/user/{userID}", handler.UserGetHandler,
	},
	Route{
		"UserAdd", "PUT", "/user", handler.UserPutHandler,
	},
}
