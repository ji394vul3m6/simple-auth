package main

import (
	"litttlebear/simple-auth/dao"
	"litttlebear/simple-auth/data"
)

var useDB dao.DB

func setDB(db dao.DB) {
	useDB = db
}

func getEnterprises() (*data.Enterprises, string) {
	enterprises, err := useDB.GetEnterprises()
	if err != nil {
		return nil, err.Error()
	}
	return enterprises, ""
}
func getEnterprise(enterpriseID string) (*data.Enterprise, string) {
	enterprise, err := useDB.GetEnterprise(enterpriseID)
	if err != nil {
		return nil, err.Error()
	}
	apps, err := useDB.GetApps(enterpriseID)
	if err != nil {
		return nil, err.Error()
	}
	adminUser, err := useDB.GetAdminUser(enterpriseID)
	if err != nil {
		return nil, err.Error()
	}

	enterprise.Apps = apps
	enterprise.AdminUser = adminUser
	return enterprise, ""
}
func getUsers(enterpriseID string) (*data.Users, string) {
	enterprises, err := useDB.GetUsers(enterpriseID)
	if err != nil {
		return nil, err.Error()
	}
	return enterprises, ""
}
func getUser(enterpriseID string, userID string) (*data.User, string) {
	user, err := useDB.GetUser(enterpriseID, userID)
	if err != nil {
		return nil, err.Error()
	}
	return user, ""
}
func getApps(enterpriseID string) (*data.Apps, string) {
	apps, err := useDB.GetApps(enterpriseID)
	if err != nil {
		return nil, err.Error()
	}
	return apps, ""
}
func getApp(enterpriseID string, appID string) (*data.App, string) {
	app, err := useDB.GetApp(enterpriseID, appID)
	if err != nil {
		return nil, err.Error()
	}
	return app, ""
}

func login(email string, passwd string) (*data.Enterprise, *data.User, string) {
	enterpriseID, user, err := useDB.GetAuthUser(email, passwd)
	if err != nil {
		return nil, nil, err.Error()
	}

	enterprise, err := useDB.GetEnterprise(enterpriseID)
	if err != nil {
		return nil, nil, err.Error()
	}
	return enterprise, user, ""
}
