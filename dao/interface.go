package dao

import "github.com/ji394vul3m6/simple-auth/data"

// DB define interface for different dao modules
type DB interface {
	GetEnterprises() (*data.Enterprises, error)
	GetEnterprise(enterpriseID string) (*data.Enterprise, error)

	GetUsers(enterpriseID string) (*data.Users, error)
	GetUser(enterpriseID string, userID string) (*data.User, error)
	GetAdminUser(enterpriseID string) (*data.User, error)
	GetAuthUser(email string, passwd string) (user *data.User, err error)

	AddUser(enterpriseID string, user *data.User) (userID string, err error)
	UpdateUser(enterpriseID string, user *data.User) error
	DeleteUser(enterpriseID string, userID string) (bool, error)

	GetRoles(enterpriseID string) ([]*data.Role, error)
	GetRole(enterpriseID string, roleID string) (*data.Role, error)
	AddRole(enterprise string, role *data.Role) (string, error)
	UpdateRole(enterprise string, roleID string, role *data.Role) (bool, error)
	DeleteRole(enterprise string, roleID string) (bool, error)
	GetUsersOfRole(enterpriseID string, roleID string) (*data.Users, error)

	GetModules(enterpriseID string) ([]*data.Module, error)

	// TODO
	DisableUser(enterpriseID string, userID string) (bool, error)

	AddEnterprise(enterprise *data.Enterprise) (string, error)
	DeleteEnterprise(enterpriseID string) (bool, error)

	GetApps(enterpriseID string) (*data.Apps, error)
	GetApp(enterpriseID string, AppID string) (*data.App, error)
	AddApp(enterpriseID string, app data.App) (*data.App, error)
	UpdateApp(enterpriseID string, app data.App) (*data.App, error)
	DisableApp(enterpriseID string, AppID string) (bool, error)
	DeleteApp(enterpriseID string, AppID string) (bool, error)
}
