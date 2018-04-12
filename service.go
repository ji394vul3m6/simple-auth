package main

import (
	"errors"

	"github.com/ji394vul3m6/simple-auth/dao"
	"github.com/ji394vul3m6/simple-auth/data"
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
	users, err := useDB.GetUsers(enterpriseID)
	if err != nil {
		return nil, err
	}
	return users, nil
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

func getRoles(enterpriseID string) ([]*data.Role, error) {
	ret, err := useDB.GetRoles(enterpriseID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func getRole(enterpriseID string, roleID string) (*data.Role, error) {
	ret, err := useDB.GetRole(enterpriseID, roleID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func deleteRole(enterpriseID string, roleID string) (bool, error) {
	ret, err := useDB.DeleteRole(enterpriseID, roleID)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func addRole(enterpriseID string, role *data.Role) (string, error) {
	return useDB.AddRole(enterpriseID, role)
}

func updateRole(enterpriseID string, roleID string, role *data.Role) (bool, error) {
	roles, err := useDB.GetUsersOfRole(enterpriseID, roleID)
	if err != nil {
		return false, err
	}
	if roles != nil && len(*roles) > 0 {
		return false, errors.New("Cannot remove role having user")
	}
	return useDB.UpdateRole(enterpriseID, roleID, role)
}

func getModules(enterpriseID string) ([]*data.Module, error) {
	ret, err := useDB.GetModules(enterpriseID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
