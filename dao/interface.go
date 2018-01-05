package dao

import "litttlebear/simple-auth/data"

// DB define interface for different dao modules
type DB interface {
	GetEnterprises() (*data.Enterprises, error)
	GetEnterprise(enterpriseID string) (*data.Enterprise, error)
	AddEnterprise(enterprise data.Enterprise) (*data.Enterprise, error)
	DeleteEnterprise(enterpriseID string) (bool, error)

	GetUsers(enterpriseID string) (*data.Users, error)
	GetUser(enterpriseID string, userID string) (*data.User, error)
	AddUser(enterpriseID string, user data.User) (*data.User, error)
	UpdateUser(enterpriseID string, user data.User) (*data.User, error)
	DisableUser(enterpriseID string, userID string) (bool, error)
	DeleteUser(enterpriseID string, userID string) (bool, error)

	GetApps(enterpriseID string) (*data.Apps, error)
	GetApp(enterpriseID string, AppID string) (*data.App, error)
	AddApp(enterpriseID string, app data.App) (*data.App, error)
	UpdateApp(enterpriseID string, app data.App) (*data.App, error)
	DisableApp(enterpriseID string, AppID string) (bool, error)
	DeleteApp(enterpriseID string, AppID string) (bool, error)
}
