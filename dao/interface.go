package dao

import "litttlebear/simple-auth/data"

type db interface {
	GetEnterprises() *data.Enterprises
	GetEnterprise() *data.Enterprise
	AddEnterprise(enterprise data.Enterprise) *data.Enterprise
	DeleteEnterprise(enterpriseID string) bool

	GetUsers(enterpriseID string) *data.Users
	GetUser(enterpriseID string, userID string) *data.User
	AddUser(enterpriseID string, user data.User) *data.User
	UpdateUser(enterpriseID string, user data.User) *data.User
	Disableuser(enterpriseID string, userID string) bool
	DeleteUser(enterpriseID string, userID string) bool

	GetApps(enterpriseID string) *data.Apps
	GetApp(enterpriseID string, AppID string) *data.App
	AddApp(enterpriseID string, app data.App) *data.App
	UpdateApp(enterpriseID string, app data.App) *data.App
	DisableApp(enterpriseID string, AppID string) bool
}
