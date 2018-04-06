package main

import (
	"encoding/json"
	"fmt"
	"litttlebear/simple-auth/data"
	"litttlebear/simple-auth/util"
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
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}

	retData, errMsg := getEnterprise(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func UsersGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}

	retData, errMsg := getUsers(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}

	userID := vars["userID"]
	if !util.IsValidUUID(userID) {
		returnBadRequest(w, "userID")
		return
	}

	retData, errMsg := getUser(enterpriseID, userID)
	returnMsg(w, errMsg, retData)
}

func AppsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}

	retData, errMsg := getApps(enterpriseID)
	returnMsg(w, errMsg, retData)
}

func AppGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}

	appID := vars["appID"]
	if !util.IsValidUUID(appID) {
		returnBadRequest(w, "appID")
		return
	}
	retData, errMsg := getApp(enterpriseID, appID)
	returnMsg(w, errMsg, retData)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	passwd := r.Form.Get("passwd")
	if !util.IsValidString(&email) || !util.IsValidString(&passwd) {
		returnBadRequest(w, "")
		return
	}

	enterprise, user, errMsg := login(email, passwd)
	if errMsg != "" {
		returnInternalError(w, errMsg)
		return
	} else if enterprise == nil && user == nil {
		returnForbidden(w)
		return
	}

	token, err := user.GenerateToken()
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}

	loginRet := data.LoginInfo{
		Token: token,
		Info:  user,
	}
	returnOKMsg(w, errMsg, loginRet)

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
		writeErrJSON(w, errMsg)
	} else {
		returnSuccess(w, retData)
	}
}

func returnNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	writeErrJSON(w, "Resource not found")
}

func returnBadRequest(w http.ResponseWriter, column string) {
	errMsg := ""
	w.WriteHeader(http.StatusBadRequest)
	if column != "" {
		errMsg = fmt.Sprintf("Column input error: %s", column)
	} else {
		errMsg = "Bad request"
	}
	writeErrJSON(w, errMsg)
}

func returnUnauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func returnForbidden(w http.ResponseWriter) {
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func returnInternalError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	writeErrJSON(w, errMsg)
}

func returnSuccess(w http.ResponseWriter, retData interface{}) {
	ret := data.Return{
		ReturnMessage: "success",
		ReturnObj:     &retData,
	}

	writeResponseJSON(w, &ret)
}

func writeErrJSON(w http.ResponseWriter, errMsg string) {
	ret := data.Return{
		ReturnMessage: errMsg,
		ReturnObj:     nil,
	}
	writeResponseJSON(w, &ret)
}

func writeResponseJSON(w http.ResponseWriter, ret *data.Return) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ret)
}
