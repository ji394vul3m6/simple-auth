package main

import (
	"encoding/json"
	"litttlebear/simple-auth/data"
	"net/http"
)

func EnterprisesGetHandler(w http.ResponseWriter, r *http.Request) {
	enterprises := getEnterprises()

	ret := data.Return{
		ReturnCode:    0,
		ReturnMessage: "success",
		ReturnObj:     &enterprises,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ret)
}
