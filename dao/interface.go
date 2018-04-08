package dao

import "litttlebear/simple-auth/data"

// DB define interface for different dao modules
type DB interface {
	GetEnterprises() (*data.Enterprises, error)
	GetEnterprise(enterpriseID string) (*data.Enterprise, error)

	AddEnterprise(enterprise *data.Enterprise) (string, error)
	DeleteEnterprise(enterpriseID string) (bool, error)

	GetUsers(enterpriseID string) (*data.Users, error)
	GetUser(enterpriseID string, userID string) (*data.User, error)
	GetAdminUser(enterpriseID string) (*data.User, error)
	GetAuthUser(email string, passwd string) (user *data.User, err error)

	AddUser(enterpriseID string, user *data.User) (userID string, err error)
	UpdateUser(enterpriseID string, user *data.User) error
	DisableUser(enterpriseID string, userID string) (bool, error)
	DeleteUser(enterpriseID string, userID string) (bool, error)

	GetApps(enterpriseID string) (*data.Apps, error)
	GetApp(enterpriseID string, AppID string) (*data.App, error)
	AddApp(enterpriseID string, app data.App) (*data.App, error)
	UpdateApp(enterpriseID string, app data.App) (*data.App, error)
	DisableApp(enterpriseID string, AppID string) (bool, error)
	DeleteApp(enterpriseID string, AppID string) (bool, error)
}
