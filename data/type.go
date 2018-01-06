package data

import (
	"encoding/json"
	"litttlebear/simple-auth/util"
)

const (
	userActive   = 1
	userInactive = 0
)

// User store the basic logging information of user
type User struct {
	ID          string  `json:"id"`
	DisplayName *string `json:"display_name"`
	Email       *string `json:"email"`
	Enterprise  *string `json:"enterprise"`
	Type        *int    `json:"type"`
	Password    *string `json:"-"`
	Status      *int    `json:"status"`
}

// Users means []User
type Users []User

// IsValid will check valid of not
func (user User) IsValid() bool {
	return util.IsValidString(user.Email) &&
		util.IsValidString(user.Password) &&
		util.IsValidMD5(*user.Password) &&
		util.IsValidString(user.Enterprise) &&
		util.IsValidUUID(*user.Enterprise)
}

// IsActive will check user is active or not
func (user User) IsActive() bool {
	return user.Status != nil && *user.Status == userActive
}

// GenerateToken will generate json web token for current user
func (user User) GenerateToken() (string, error) {
	return util.GetJWTTokenWithCustomInfo(&user)
}

// SetValueWithToken will return an userObj from custom column of token
func (user *User) SetValueWithToken(tokenString string) error {
	info, err := util.ResolveJWTToken(tokenString)
	if err != nil {
		return err
	}
	jsonByte, _ := json.Marshal(info)

	userInfo := User{}
	err = json.Unmarshal(jsonByte, &userInfo)
	if err != nil {
		return err
	}
	user.CopyValue(userInfo)
	return nil
}

func (user *User) CopyValue(source User) {
	user.ID = source.ID
	user.DisplayName = source.DisplayName
	user.Email = source.Email
	user.Enterprise = source.Enterprise
	user.Type = source.Type
	user.Password = source.Password
	user.Status = source.Status
}

// App store basic app usage information in it
type App struct {
	ID     string  `json:"id"`
	Name   *string `json:"name"`
	Status *int    `json:"status"`
	// ValidStart and ValidEnd will store in timestamp format
	ValidStart *int `json:"valid_start"`
	ValidEnd   *int `json:"valid_end"`
	ValidCount *int `json:"valid_count"`
	// Enterprise *string `json:"enterprise"`
}

// Apps is array of App
type Apps []App

// Enterprise store the basic logging information of enterprise
// which can has multi app and user in it
type Enterprise struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	AdminUser *User   `json:"admin_user,omitempty"`
	Apps      *Apps   `json:"apps,omitempty"`
}

// Enterprises is array of Enterprise
type Enterprises []Enterprise

// LoginInfo is struct return when calling login
type LoginInfo struct {
	Token string `json:"token"`
	Info  *User  `json:"info"`
}
