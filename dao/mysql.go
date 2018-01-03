package dao

import (
	"database/sql"
	"fmt"
	"litttlebear/simple-auth/data"
	"log"
)

var connectDB *sql.DB

type mySQLController struct{}

func initDB(host string, port int, dbName string, account string, password string) bool {
	var dbString string
	if port == 0 {
		dbString = fmt.Sprintf("%s:%s@%s/%s", account, password, host, dbName)
	} else {
		dbString = fmt.Sprintf("%s:%s@%s:%d/%s", account, password, host, port, dbName)
	}
	db, err := sql.Open("mysql", dbString)

	if err != nil {
		log.Printf("Connect to db[%s] fail\n", dbString)
		return false
	}

	connectDB = db
	return true
}

func checkDB() bool {
	if connectDB == nil {
		log.Fatal("connectDB is nil, db is !initialized properly")
		return false
	}
	connectDB.Ping()
	return true
}

func GetEnterprise() *data.Enterprise {
	if !checkDB() {
		return nil
	}
	return nil
}
func AddEnterprise(enterprise data.Enterprise) *data.Enterprise {
	if !checkDB() {
		return nil
	}
	return nil
}

func DeleteEnterprise(enterpriseID string) bool {
	return true
}

func GetUsers(enterpriseID string) *data.Users {
	if !checkDB() {
		return nil
	}
	return nil
}
func GetUser(enterpriseID string, userID string) *data.User {
	if !checkDB() {
		return nil
	}
	return nil
}
func AddUser(enterpriseID string, user data.User) *data.User {
	if !checkDB() {
		return nil
	}
	return nil
}
func UpdateUser(enterpriseID string, user data.User) *data.User {
	if !checkDB() {
		return nil
	}
	return nil
}
func DeleteUser(enterpriseID string, userID string) bool {
	return false
}

func GetApps(enterpriseID string) *data.Apps {
	if !checkDB() {
		return nil
	}
	return nil
}
func GetApp(enterpriseID string, AppID string) *data.App {
	if !checkDB() {
		return nil
	}
	return nil
}
func AddApp(enterpriseID string, app data.App) *data.App {
	if !checkDB() {
		return nil
	}
	return nil
}
func UpdateApp(enterpriseID string, app data.App) *data.App {
	if !checkDB() {
		return nil
	}
	return nil
}
func DeleteApp(enterpriseID string, AppID string) bool {
	return true
}
