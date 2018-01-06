package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"litttlebear/simple-auth/dao"
	"litttlebear/simple-auth/data"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	AuthPass = false
	AuthFail = true

	URLPrefix = "/auth"
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
		Route{
			"GetEnterprises", "GET", "/enterprise/{enterpriseID}", EnterpriseGetHandler,
		},
		Route{
			"GetUsers", "GET", "/enterprise/{enterpriseID}/users", UsersGetHandler,
		},
		Route{
			"GetUser", "GET", "/enterprise/{enterpriseID}/user/{userID}", UserGetHandler,
		},
		Route{
			"GetApps", "GET", "/enterprise/{enterpriseID}/apps", AppsGetHandler,
		},
		Route{
			"GetApp", "GET", "/enterprise/{enterpriseID}/app/{appID}", AppGetHandler,
		},
		Route{
			"Login", "POST", "/login", LoginHandler,
		},
	}
	return &routes
}

func setUpDB() {
	db := dao.MYSQLController{}
	db.InitDB("127.0.0.1", 3306, "auth", "root", "password")
	setDB(&db)
}

func checkAuth(r *http.Request, rm *mux.RouteMatch) bool {
	log.Printf("Access: %s %s", r.Method, r.RequestURI)
	if r.RequestURI == URLPrefix+"/login" {
		return AuthPass
	}

	authorization := r.Header.Get("Authorization")
	vals := strings.Split(authorization, " ")
	if len(vals) < 2 {
		return AuthFail
	}

	userInfo := data.User{}
	err := userInfo.SetValueWithToken(vals[1])
	if err != nil {
		return AuthFail
	}
	return AuthPass
}

func main() {
	routes := setUpRoutes()
	setUpDB()

	router := mux.NewRouter().StrictSlash(true)

	// Use matcher func first to check for all request first.
	// It will be used for auth check in future commit
	router.
		MatcherFunc(checkAuth).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})

	for _, route := range *routes {
		router.
			Methods(route.Method).
			Path(URLPrefix + route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
		log.Printf("Setup for path [%s:%s]", route.Method, URLPrefix+route.Pattern)
	}

	log.Printf("Start server on port %d", 11180)
	log.Fatal(http.ListenAndServe(":11180", router))
}
