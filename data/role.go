package data

type Role struct {
	UUID        string              `json:"id"`
	Name        string              `json:"name"`
	Discription string              `json:"description"`
	Privileges  map[string][]string `json:"privileges"`
}

type Module struct {
	ID       int      `json:"-"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}
