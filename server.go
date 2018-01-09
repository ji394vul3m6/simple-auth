package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"litttlebear/simple-auth/dao"
	"litttlebear/simple-auth/data"
	"litttlebear/simple-auth/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	prefixURL = "/auth"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	GrantType   []interface{}
}

type Routes []Route

var routes Routes

func setUpRoutes() {
	routes = Routes{
		Route{
			"Base", "GET", "/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			}, []interface{}{0, 1, 2},
		},
		Route{
			"GetEnterprises", "GET", "/enterprises", EnterprisesGetHandler, []interface{}{0},
		},
		Route{
			"GetEnterprises", "GET", "/enterprise/{enterpriseID}", EnterpriseGetHandler, []interface{}{0},
		},
		Route{
			"GetUsers", "GET", "/enterprise/{enterpriseID}/users", UsersGetHandler, []interface{}{0, 1},
		},
		Route{
			"GetUser", "GET", "/enterprise/{enterpriseID}/user/{userID}", UserGetHandler, []interface{}{0, 1},
		},
		Route{
			"GetApps", "GET", "/enterprise/{enterpriseID}/apps", AppsGetHandler, []interface{}{0, 1},
		},
		Route{
			"GetApp", "GET", "/enterprise/{enterpriseID}/app/{appID}", AppGetHandler, []interface{}{0, 1},
		},
		Route{
			"Login", "POST", "/login", LoginHandler, []interface{}{},
		},
	}
}

func setUpDB() {
	db := dao.MYSQLController{}
	db.InitDB("127.0.0.1", 3306, "auth", "root", "password")
	setDB(&db)
}

func checkAuth(r *http.Request, route Route) bool {
	log.Printf("Access: %s %s", r.Method, r.RequestURI)
	if len(route.GrantType) == 0 {
		log.Print("[Auth check] pass: no need")
		return true
	}

	authorization := r.Header.Get("Authorization")
	vals := strings.Split(authorization, " ")
	if len(vals) < 2 {
		log.Print("[Auth check] Auth fail: no header")
		return false
	}

	userInfo := data.User{}
	err := userInfo.SetValueWithToken(vals[1])
	if err != nil {
		log.Printf("[Auth check] Auth fail: no valid token [%s]", err.Error())
		return false
	}

	if !util.IsInSlice(*userInfo.Type, route.GrantType) {
		log.Printf("[Auth check] Need user be [%#v], get [%d]", route.GrantType, *userInfo.Type)
		return false
	}

	// Type 1 can only check enterprise of itself
	if *userInfo.Type == 1 {
		vars := mux.Vars(r)
		enterpriseID := vars["enterpriseID"]
		if enterpriseID != *userInfo.Enterprise {
			log.Printf("[Auth check] admin of [%s] can not access [%s]", *userInfo.Enterprise, enterpriseID)
			return false
		}
	}

	return true
}

func main() {
	setUpRoutes()
	setUpDB()

	router := mux.NewRouter().StrictSlash(true)

	for idx := range routes {
		route := routes[idx]
		router.
			Methods(route.Method).
			Path(prefixURL + route.Pattern).
			Name(route.Name).
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if checkAuth(r, route) {
					route.HandlerFunc(w, r)
				} else {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
				}
			})
		log.Printf("Setup for path [%s:%s]", route.Method, prefixURL+route.Pattern)
	}

	log.Printf("Start server on port %d", 11180)
	log.Fatal(http.ListenAndServe(":11180", router))
}
