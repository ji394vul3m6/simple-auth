package dao

import (
	"database/sql"
	"fmt"
	"litttlebear/simple-auth/data"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

const (
	enterpriseTable = "enterprises"
	userTable       = "users"
	appTable        = "apps"
)

type MYSQLController struct {
	connectDB *sql.DB
}

func (controller *MYSQLController) InitDB(host string, port int, dbName string, account string, password string) bool {
	var dbString string
	if port == 0 {
		dbString = fmt.Sprintf("%s:%s@%s/%s", account, password, host, dbName)
	} else {
		dbString = fmt.Sprintf("%s:%s@(%s:%d)/%s", account, password, host, port, dbName)
	}
	log.Printf("Connect to db [%s]", dbString)
	db, err := sql.Open("mysql", dbString)

	if err != nil {
		log.Printf("Connect to db[%s] fail: [%s]\n", dbString, err.Error())
		return false
	}

	controller.connectDB = db
	return true
}

func (controller MYSQLController) checkDB() bool {
	if controller.connectDB == nil {
		log.Fatal("connectDB is nil, db is !initialized properly")
		return false
	}
	controller.connectDB.Ping()
	return true
}

func (controller MYSQLController) GetEnterprises() *data.Enterprises {
	if !controller.checkDB() {
		return nil
	}
	enterprises := make(data.Enterprises, 0)
	rows, err := controller.connectDB.Query(fmt.Sprintf("SELECT uuid,name from %s", userTable))
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		enterprise := data.Enterprise{}
		err := rows.Scan(&enterprise.ID, &enterprise.Name)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil
		}
		enterprises = append(enterprises, enterprise)
	}

	return &enterprises
}
func (controller MYSQLController) GetEnterprise() *data.Enterprise {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) AddEnterprise(enterprise data.Enterprise) *data.Enterprise {
	if !controller.checkDB() {
		return nil
	}
	return nil
}

func (controller MYSQLController) DeleteEnterprise(enterpriseID string) bool {
	return true
}

func (controller MYSQLController) GetUsers(enterpriseID string) *data.Users {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) GetUser(enterpriseID string, userID string) *data.User {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) AddUser(enterpriseID string, user data.User) *data.User {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) UpdateUser(enterpriseID string, user data.User) *data.User {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) DisableUser(enterpriseID string, userID string) bool {
	return false
}
func (controller MYSQLController) DeleteUser(enterpriseID string, userID string) bool {
	return false
}

func (controller MYSQLController) GetApps(enterpriseID string) *data.Apps {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) GetApp(enterpriseID string, AppID string) *data.App {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) AddApp(enterpriseID string, app data.App) *data.App {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) UpdateApp(enterpriseID string, app data.App) *data.App {
	if !controller.checkDB() {
		return nil
	}
	return nil
}
func (controller MYSQLController) DisableApp(enterpriseID string, AppID string) bool {
	return true
}
func (controller MYSQLController) DeleteApp(enterpriseID string, AppID string) bool {
	return true
}
