package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"litttlebear/simple-auth/dao"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func setUpRoutes() *Routes {
	var routes = Routes{
		Route{
			"Base", "GET", "/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			},
		},
		Route{
			"GetEnterprises", "GET", "/enterprises", EnterprisesGetHandler,
		},
	}
	return &routes
}

func setUpDB() {
	db := dao.MYSQLController{}
	db.InitDB("127.0.0.1", 3306, "auth", "root", "password")
	setDB(&db)
}

func main() {
	routes := setUpRoutes()
	setUpDB()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		log.Printf("Setup for path [%s:%s]", route.Method, route.Pattern)
	}

	log.Printf("Start server on port %d", 11180)
	log.Fatal(http.ListenAndServe(":11180", router))
}
