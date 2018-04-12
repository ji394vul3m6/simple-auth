package data

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
