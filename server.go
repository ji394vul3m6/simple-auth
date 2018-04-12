package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/ji394vul3m6/simple-auth/data"
	"github.com/ji394vul3m6/simple-auth/util"

	"github.com/ji394vul3m6/simple-auth/dao"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	prefixURL = "/auth"
)

// Route define the end point of server
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc

	// 0 means super admin can use this API
	// 1 means enterprise admin can use this API
	// 2 means user in enterprise can use this API
	GrantType []interface{}
}

type Routes []Route

var routes Routes

func setUpRoutes() {
	routes = Routes{
		Route{"GetEnterprises", "GET", "/enterprises", EnterprisesGetHandler, []interface{}{0}},
		Route{"GetEnterprises", "GET", "/enterprise/{enterpriseID}", EnterpriseGetHandler, []interface{}{0, 1, 2}},
		Route{"GetUsers", "GET", "/enterprise/{enterpriseID}/users", UsersGetHandler, []interface{}{0, 1}},
		Route{"GetUser", "GET", "/enterprise/{enterpriseID}/user/{userID}", UserGetHandler, []interface{}{0, 1, 2}},
		Route{"GetApps", "GET", "/enterprise/{enterpriseID}/apps", AppsGetHandler, []interface{}{0, 1, 2}},
		Route{"GetApp", "GET", "/enterprise/{enterpriseID}/app/{appID}", AppGetHandler, []interface{}{0, 1, 2}},
		Route{"Login", "POST", "/login", LoginHandler, []interface{}{}},
		Route{"ValidateToken", "GET", "/token/{token}", ValidateTokenHandler, []interface{}{}},

		Route{"AddUser", "POST", "/enterprise/{enterpriseID}/user", UserAddHandler, []interface{}{0, 1, 2}},
		Route{"UpdateUser", "PUT", "/enterprise/{enterpriseID}/user/{userID}", UserUpdateHandler, []interface{}{0, 1, 2}},
		Route{"DeleteUser", "DELETE", "/enterprise/{enterpriseID}/user/{userID}", UserDeleteHandler, []interface{}{0, 1, 2}},

		Route{"GetRoles", "GET", "/enterprise/{enterpriseID}/roles", RolesGetHandler, []interface{}{0, 1, 2}},
		Route{"GetRole", "GET", "/enterprise/{enterpriseID}/role/{roleID}", RoleGetHandler, []interface{}{0, 1, 2}},
		Route{"AddRole", "POST", "/enterprise/{enterpriseID}/role", RoleAddHandler, []interface{}{0, 1, 2}},
		Route{"UpdateRole", "PUT", "/enterprise/{enterpriseID}/role/{roleID}", RoleUpdateHandler, []interface{}{0, 1, 2}},
		Route{"DeleteRole", "DELETE", "/enterprise/{enterpriseID}/role/{roleID}", RoleDeleteHandler, []interface{}{0, 1, 2}},
		Route{"GetModules", "GET", "/enterprise/{enterpriseID}/modules", ModulesGetHandler, []interface{}{0, 1, 2}},

		// Route{"AddModules", "GET", "/enterprise/{enterpriseID}/module", ModuleAddHandler, []interface{}{0}},
		// Route{"UpdateModules", "GET", "/enterprise/{enterpriseID}/module/{moduleCode}", ModuleUpdateHandler, []interface{}{0}},
		// Route{"DeleteModules", "GET", "/enterprise/{enterpriseID}/module/{moduleCode}", ModuleDeleteHandler, []interface{}{0}},
		// Route{"AddApp", "GET", "/enterprise/{enterpriseID}/app", AppAddHandler, []interface{}{0, 1}},
		// Route{"UpdateApp", "GET", "/enterprise/{enterpriseID}/app/{appID}", AppUpdateHandler, []interface{}{0, 1}},
		// Route{"DeleteApp", "GET", "/enterprise/{enterpriseID}/app/{appID}", AppDeleteHandler, []interface{}{0, 1}},
	}
}

func setUpDB() {
	db := dao.MYSQLController{}
	url, port, user, passwd, dbName := util.GetMySQLConfig()
	log.Printf("Init mysql: %s:%s@%s:%d/%s\n", user, passwd, url, port, dbName)
	db.InitDB(url, port, dbName, user, passwd)
	setDB(&db)
}

func checkAuth(r *http.Request, route Route) bool {
	log.Printf("Access: %s %s", r.Method, r.RequestURI)
	if len(route.GrantType) == 0 {
		log.Println("[Auth check] pass: no need")
		return true
	}

	authorization := r.Header.Get("Authorization")
	vals := strings.Split(authorization, " ")
	if len(vals) < 2 {
		log.Println("[Auth check] Auth fail: no header")
		return false
	}

	userInfo := data.User{}
	err := userInfo.SetValueWithToken(vals[1])
	if err != nil {
		log.Printf("[Auth check] Auth fail: no valid token [%s]\n", err.Error())
		return false
	}

	if !util.IsInSlice(userInfo.Type, route.GrantType) {
		log.Printf("[Auth check] Need user be [%v], get [%d]\n", route.GrantType, userInfo.Type)
		return false
	}

	vars := mux.Vars(r)
	// Type 1 can only check enterprise of itself
	// Type 2 can only check enterprise of itself and user info of itself
	if userInfo.Type == 1 || userInfo.Type == 2 {
		enterpriseID := vars["enterpriseID"]
		if enterpriseID != *userInfo.Enterprise {
			log.Printf("[Auth check] user of [%s] can not access [%s]\n", *userInfo.Enterprise, enterpriseID)
			return false
		}
	}

	if userInfo.Type == 2 {
		userID := vars["userID"]
		if userID != "" && userID != userInfo.ID {
			log.Printf("[Auth check] user [%s] can not access other users' info\n", userInfo.ID)
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
					if route.HandlerFunc != nil {
						route.HandlerFunc(w, r)
					}
				} else {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
				}
			})
		log.Printf("Setup for path [%s:%s], %+v", route.Method, prefixURL+route.Pattern, route.GrantType)
	}

	log.Printf("Start server on port %d", 11180)
	log.Fatal(http.ListenAndServe(":11180", router))
}

func getRequester(r *http.Request) *data.User {
	authorization := r.Header.Get("Authorization")
	vals := strings.Split(authorization, " ")
	if len(vals) < 2 {
		return nil
	}

	userInfo := data.User{}
	err := userInfo.SetValueWithToken(vals[1])
	if err != nil {
		return nil
	}

	return &userInfo
}
