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
func getUsers(enterpriseID string) (*data.Users, error) {
	enterprises, err := useDB.GetUsers(enterpriseID)
	if err != nil {
		return nil, err
	}
	return enterprises, nil
}
func getUser(enterpriseID string, userID string) (*data.User, error) {
	user, err := useDB.GetUser(enterpriseID, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
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
	user, err := useDB.GetAuthUser(email, passwd)
	if err != nil {
		return nil, nil, err.Error()
	}

	var enterprise *data.Enterprise
	if user.Enterprise != nil {
		enterprise, err = useDB.GetEnterprise(*user.Enterprise)
		if err != nil {
			return nil, nil, err.Error()
		}
	}
	return enterprise, user, ""
}

func addUser(enterpriseID string, user *data.User) (string, error) {
	return useDB.AddUser(enterpriseID, user)
}

func deleteUser(enterpriseID string, userID string) error {
	_, err := useDB.DeleteUser(enterpriseID, userID)
	return err
}

func updateUser(enterpriseID string, user *data.User) error {
	return useDB.UpdateUser(enterpriseID, user)
}
