package main

import (
	"litttlebear/simple-auth/dao"
	"litttlebear/simple-auth/data"
)

var useDB dao.DB

func setDB(db dao.DB) {
	useDB = db
}

func getEnterprises() *data.Enterprises {
	enterprises := useDB.GetEnterprises()
	return enterprises
}
