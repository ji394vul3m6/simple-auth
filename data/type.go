package data

import (
	"litttlebear/simple-auth/util"
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
// User is valid only if username, email and password are not empty
func (user User) IsValid() bool {
	return util.IsValidString(user.Email) && util.IsValidString(user.Password)
}

// GenerateToken will generate json web token for current user
func (user User) GenerateToken() string {
	return "temp_token"
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
	AdminUser *User   `json:"admin_user"`
	Apps      *Apps   `json:"apps"`
}

// Enterprises is array of Enterprise
type Enterprises []Enterprise

// LoginInfo is struct return when calling login
type LoginInfo struct {
	Token string `json:"token"`
	Info  *User  `json:"info"`
}
