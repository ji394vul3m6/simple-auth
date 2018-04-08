package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"litttlebear/simple-auth/data"
	"litttlebear/simple-auth/enum"
	"litttlebear/simple-auth/util"
	"net/http"
	"reflect"
	"strconv"

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

	retData, err := getUsers(enterpriseID)
	if err != nil {
		returnInternalError(w, err.Error())
	} else {
		returnSuccess(w, retData)
	}
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

	retData, err := getUser(enterpriseID, userID)
	if err != nil {
		returnInternalError(w, err.Error())
	} else {
		returnSuccess(w, retData)
	}
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

func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	requester := getRequester(r)
	vars := mux.Vars(r)
	enterpriseID := vars["enterpriseID"]
	if !util.IsValidUUID(enterpriseID) {
		returnBadRequest(w, "enterpriseID")
		return
	}
	user, err := parseAddUserFromRequest(r)
	if err != nil {
		returnBadRequest(w, err.Error())
		return
	}

	if requester.Type > user.Type {
		returnForbidden(w)
		return
	}

	id, err := addUser(enterpriseID, user)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}
	newUser, err := getUser(enterpriseID, id)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}
	returnSuccess(w, newUser)
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	requester := getRequester(r)
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

	user, err := getUser(enterpriseID, userID)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	} else if user == nil {
		returnSuccess(w, "")
		return
	}

	if requester.Type > user.Type {
		returnForbidden(w)
		return
	}

	err = deleteUser(enterpriseID, userID)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}

	returnSuccess(w, "")
}

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	origUser, err := getUser(enterpriseID, userID)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	} else if origUser == nil {
		returnNotFound(w)
		return
	}

	newUser, err := parseUpdateUserFromRequest(r)
	if err != nil {
		returnBadRequest(w, err.Error())
		return
	}

	newUser.ID = userID
	newUser.Enterprise = &enterpriseID
	err = updateUser(enterpriseID, newUser)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}

	updatedUser, err := getUser(enterpriseID, userID)
	if err != nil {
		returnInternalError(w, err.Error())
		return
	}
	returnSuccess(w, updatedUser)
}

func loadUserFromRequest(r *http.Request) *data.User {
	user := data.User{}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user.Email = &email
	user.DisplayName = &name
	user.Password = &password

	userType, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		userType = enum.NormalUser
	} else if userType > enum.NormalUser || userType < enum.AdminUser {
		userType = enum.NormalUser
	}
	user.Type = userType

	return &user
}
func parseAddUserFromRequest(r *http.Request) (*data.User, error) {
	user := loadUserFromRequest(r)

	if user.Email == nil || *user.Email == "" {
		return nil, errors.New("invalid email")
	}
	if user.Password == nil || *user.Password == "" {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
func parseUpdateUserFromRequest(r *http.Request) (*data.User, error) {
	user := loadUserFromRequest(r)

	if user.Email == nil || *user.Email == "" {
		return nil, errors.New("invalid email")
	}

	return user, nil
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
