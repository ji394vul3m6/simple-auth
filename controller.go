package main

import (
	"encoding/json"
	"fmt"
	"litttlebear/simple-auth/data"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func EnterprisesGetHandler(w http.ResponseWriter, r *http.Request) {
	retData, errMsg := getEnterprises()
	returnOKMsg(w, errMsg, retData)
}

func EnterpriseGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	retData, errMsg := getEnterprise(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func UsersGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	retData, errMsg := getUsers(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	userID := vars["userID"]
	retData, errMsg := getUser(enterpriseID, userID)
	returnMsg(w, errMsg, retData)
}

func AppsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	retData, errMsg := getApps(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func AppGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	appID := vars["appID"]
	retData, errMsg := getApp(enterpriseID, appID)
	returnMsg(w, errMsg, retData)
}

func returnMsg(w http.ResponseWriter, errMsg string, retData interface{}) {
	if reflect.ValueOf(retData).IsNil() && errMsg == "" {
		returnNotFound(w)
	} else {
		returnOKMsg(w, errMsg, retData)
	}
}

func returnOKMsg(w http.ResponseWriter, errMsg string, retData interface{}) {
	if errMsg != "" {
		returnFail(w, errMsg)
	} else {
		returnSuccess(w, retData)
	}
}

func returnNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	returnFail(w, "Resource not found")
}

func returnBadRequest(w http.ResponseWriter, column string) {
	w.WriteHeader(http.StatusBadRequest)
	if column != "" {
		returnFail(w, fmt.Sprintf("Column input error: %s", column))
	} else {
		returnFail(w, "Bad request")
	}
}

func returnUnauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func returnFail(w http.ResponseWriter, errMsg string) {
	ret := data.Return{
		ReturnMessage: errMsg,
		ReturnObj:     nil,
	}

	writeResponseJSON(w, &ret)
}

func returnSuccess(w http.ResponseWriter, retData interface{}) {
	ret := data.Return{
		ReturnMessage: "success",
		ReturnObj:     &retData,
	}

	writeResponseJSON(w, &ret)
}

func writeResponseJSON(w http.ResponseWriter, ret *data.Return) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ret)
}
