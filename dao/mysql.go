package dao

import (
	"database/sql"
	"fmt"
	"litttlebear/simple-auth/data"
	"log"
	"runtime"

	"emotibot.com/emotigo/module/admin-api/util"
)

const (
	enterpriseTable = "enterprises"
	userTable       = "users"
	appTable        = "apps"
	userInfoTable   = "user_info"

	userColumnList = "uuid,display_name,email,enterprise,type,status"
)

type MYSQLController struct {
	connectDB *sql.DB
}

func (controller *MYSQLController) InitDB(host string, port int, dbName string, account string, password string) bool {
	var dbString string
	if port == 0 {
		dbString = fmt.Sprintf("%s:%s@%s/%s", account, password, host, dbName)
	} else {
		dbString = fmt.Sprintf("%s:%s@(%s:%d)/%s", account, password, host, port, dbName)
	}
	log.Printf("Connect to db [%s]", dbString)
	db, err := sql.Open("mysql", dbString)

	if err != nil {
		log.Printf("Connect to db[%s] fail: [%s]\n", dbString, err.Error())
		return false
	}

	controller.connectDB = db
	return true
}

func (controller MYSQLController) checkDB() (bool, error) {
	if controller.connectDB == nil {
		log.Fatal("connectDB is nil, db is !initialized properly")
		return false, fmt.Errorf("DB hasn't init")
	}
	controller.connectDB.Ping()
	return true, nil
}

func (controller MYSQLController) GetEnterprises() (*data.Enterprises, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	enterprises := make(data.Enterprises, 0)
	rows, err := controller.connectDB.Query(fmt.Sprintf("SELECT uuid,name from %s", enterpriseTable))
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		enterprise := data.Enterprise{}
		err := rows.Scan(&enterprise.ID, &enterprise.Name)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		enterprises = append(enterprises, enterprise)
	}

	return &enterprises, nil
}
func (controller MYSQLController) GetEnterprise(enterpriseID string) (*data.Enterprise, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	rows, err := controller.connectDB.Query(fmt.Sprintf("SELECT uuid,name from %s where uuid = ?", enterpriseTable), enterpriseID)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		enterprise := data.Enterprise{}
		err := rows.Scan(&enterprise.ID, &enterprise.Name)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		return &enterprise, nil
	}

	return nil, nil
}
func (controller MYSQLController) AddEnterprise(enterprise *data.Enterprise) (string, error) {
	ok, err := controller.checkDB()
	if !ok {
		return "", err
	}
	return "", nil
}

func (controller MYSQLController) DeleteEnterprise(enterpriseID string) (bool, error) {
	return false, nil
}

func scanSingleRowToUser(row *sql.Row) (*data.User, error) {
	user := data.User{}
	err := row.Scan(&user.ID, &user.DisplayName, &user.Email, &user.Enterprise, &user.Type, &user.Status)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func scanRowToUser(rows *sql.Rows) (*data.User, error) {
	user := data.User{}
	err := rows.Scan(&user.ID, &user.DisplayName, &user.Email, &user.Enterprise, &user.Type, &user.Status)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (controller MYSQLController) GetUsers(enterpriseID string) (*data.Users, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}

	userInfoMap, err := controller.getUsersInfo(enterpriseID)
	if err != nil {
		logDBError(err)
		return nil, err
	}

	users := make(data.Users, 0)
	rows, err := controller.connectDB.Query(fmt.Sprintf("SELECT %s from %s where enterprise = ?",
		userColumnList, userTable), enterpriseID)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := scanRowToUser(rows)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}

		if info, ok := userInfoMap[user.ID]; ok {
			user.CustomInfo = info
		}

		users = append(users, *user)
	}

	return &users, nil
}
func (controller MYSQLController) GetUser(enterpriseID string, userID string) (*data.User, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}

	queryStr := fmt.Sprintf("SELECT %s from %s where enterprise = ? and uuid = ?",
		userColumnList, userTable)
	row := controller.connectDB.QueryRow(queryStr, enterpriseID, userID)
	if err != nil {
		logDBError(err)
		return nil, err
	}
	user, err := scanSingleRowToUser(row)
	if err != nil {
		logDBError(err)
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	info, err := controller.getUserInfo(enterpriseID, userID)
	if err != nil {
		logDBError(err)
		return nil, err
	}
	user.CustomInfo = info

	return user, nil
}
func (controller MYSQLController) GetAdminUser(enterpriseID string) (*data.User, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	queryStr := fmt.Sprintf(
		"SELECT u.%s FROM %s as u LEFT JOIN %s as e ON e.admin_user = u.uuid AND e.uuid = ?",
		userColumnList, userTable, enterpriseTable)
	rows, err := controller.connectDB.Query(queryStr, enterpriseID)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanRowToUser(rows)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		return user, nil
	}

	return nil, nil
}
func (controller MYSQLController) GetAuthUser(email string, passwd string) (*data.User, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}

	queryStr := fmt.Sprintf("SELECT %s from %s where email = ? and password = ?",
		userColumnList, userTable)
	rows, err := controller.connectDB.Query(queryStr, email, passwd)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanRowToUser(rows)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		info, err := controller.getUserInfo(*user.Enterprise, user.ID)
		if err != nil {
			logDBError(err)
			return nil, err
		}
		user.CustomInfo = info
		return user, nil
	}

	return nil, nil
}
func (controller MYSQLController) AddUser(enterpriseID string, user *data.User) (string, error) {
	ok, err := controller.checkDB()
	if !ok {
		return "", err
	}

	queryStr := fmt.Sprintf("INSERT INTO %s (uuid, display_name, email, enterprise, type, password) VALUES (UUID(), ?, ?, ?, ?, ?)", userTable)
	ret, err := controller.connectDB.Exec(queryStr, user.DisplayName, user.Email, enterpriseID, user.Type, user.Password)
	if err != nil {
		return "", err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return "", err
	}

	uuid, err := controller.getUserUUID(enterpriseID, id)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
func (controller MYSQLController) getUserUUID(enterpriseID string, userID int64) (string, error) {
	ok, err := controller.checkDB()
	if !ok {
		return "", err
	}

	queryStr := fmt.Sprintf("SELECT uuid from %s WHERE enterprise = ? and id = ?", userTable)
	row := controller.connectDB.QueryRow(queryStr, enterpriseID, userID)

	ret := ""
	err = row.Scan(&ret)
	return ret, err
}

func (controller MYSQLController) UpdateUser(enterpriseID string, user *data.User) error {
	ok, err := controller.checkDB()
	if !ok {
		return err
	}
	var queryStr string
	var params []interface{}
	if user.Password == nil || *user.Password == "" {
		queryStr = fmt.Sprintf(`UPDATE %s SET
			display_name = ?, email = ?, type = ?
			WHERE uuid = ? AND enterprise = ?`, userTable)
		params = []interface{}{user.DisplayName, user.Email, user.Type, user.ID, user.Enterprise}
	} else {
		queryStr = fmt.Sprintf(`UPDATE %s SET
			display_name = ?, email = ?, type = ?,
			password = ? WHERE uuid = ? AND enterprise = ?`, userTable)
		params = []interface{}{user.DisplayName, user.Email, user.Type, user.Password, user.ID, user.Enterprise}
	}
	_, err = controller.connectDB.Exec(queryStr, params...)
	return err
}
func (controller MYSQLController) DisableUser(enterpriseID string, userID string) (bool, error) {
	return false, nil
}
func (controller MYSQLController) DeleteUser(enterpriseID string, userID string) (bool, error) {
	ok, err := controller.checkDB()
	if !ok {
		return false, err
	}

	t, err := controller.connectDB.Begin()
	if err != nil {
		return false, err
	}
	defer clearTransition(t)

	queryStr := fmt.Sprintf("DELETE FROM %s WHERE user_id = ?", userInfoTable)
	_, err = t.Exec(queryStr, userID)
	if err != nil {
		return false, err
	}

	queryStr = fmt.Sprintf("DELETE FROM %s WHERE enterprise = ? AND uuid = ?", userTable)
	_, err = t.Exec(queryStr, enterpriseID, userID)
	if err != nil {
		return false, err
	}
	t.Commit()
	return true, nil
}
func (controller MYSQLController) getUserInfo(enterpriseID string, userID string) (ret *map[string]string, err error) {
	err = nil
	ok, err := controller.checkDB()
	if !ok {
		return
	}

	queryStr := fmt.Sprintf(`SELECT col.column, info.value
		FROM user_column as col, %s as info
		WHERE info.column_id = col.id AND info.user_id = ? AND col.enterprise = ?`, userInfoTable)
	rows, err := controller.connectDB.Query(queryStr, userID, enterpriseID)
	if err != nil {
		return
	}
	defer rows.Close()

	infoMap := make(map[string]string)
	for rows.Next() {
		var key string
		var val string
		err = rows.Scan(&key, &val)
		if err != nil {
			return
		}
		infoMap[key] = val
	}
	ret = &infoMap
	return
}
func (controller MYSQLController) getUsersInfo(enterpriseID string) (ret map[string]map[string]string, err error) {
	err = nil
	ok, err := controller.checkDB()
	if !ok {
		return
	}

	queryStr := fmt.Sprintf(`SELECT info.user_id, col.column, info.value
		FROM user_column as col, %s as info
		WHERE info.column_id = col.id AND col.enterprise = ?`, userInfoTable)
	rows, err := controller.connectDB.Query(queryStr, enterpriseID)
	if err != nil {
		return
	}
	defer rows.Close()

	ret = make(map[string]map[string]string)
	for rows.Next() {
		var userID string
		var key string
		var val string
		err = rows.Scan(&userID, &key, &val)
		if err != nil {
			return
		}
		if userInfo, ok := ret[userID]; !ok {
			ret[userID] = map[string]string{
				key: val,
			}
		} else {
			userInfo[key] = val
		}
	}
	return
}

func (controller MYSQLController) GetApps(enterpriseID string) (*data.Apps, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	apps := make(data.Apps, 0)
	queryStr := fmt.Sprintf("SELECT uuid,name,UNIX_TIMESTAMP(start),UNIX_TIMESTAMP(end),UNIX_TIMESTAMP(count),status from %s where enterprise = ?", appTable)
	rows, err := controller.connectDB.Query(queryStr, enterpriseID)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		app := data.App{}
		err := rows.Scan(&app.ID, &app.Name, &app.ValidStart, &app.ValidEnd, &app.ValidCount, &app.Status)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		apps = append(apps, app)
	}

	return &apps, nil
}
func (controller MYSQLController) GetApp(enterpriseID string, AppID string) (*data.App, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	queryStr := fmt.Sprintf("SELECT uuid,name,UNIX_TIMESTAMP(start),UNIX_TIMESTAMP(end),UNIX_TIMESTAMP(count),status from %s where enterprise = ? and uuid = ?", appTable)
	rows, err := controller.connectDB.Query(queryStr, enterpriseID, AppID)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		app := data.App{}
		err := rows.Scan(&app.ID, &app.Name, &app.ValidStart, &app.ValidEnd, &app.ValidCount, &app.Status)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
			return nil, err
		}
		return &app, nil
	}

	return nil, nil
}
func (controller MYSQLController) AddApp(enterpriseID string, app data.App) (*data.App, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	return nil, nil
}
func (controller MYSQLController) UpdateApp(enterpriseID string, app data.App) (*data.App, error) {
	ok, err := controller.checkDB()
	if !ok {
		return nil, err
	}
	return nil, nil
}
func (controller MYSQLController) DisableApp(enterpriseID string, AppID string) (bool, error) {
	return false, nil
}
func (controller MYSQLController) DeleteApp(enterpriseID string, AppID string) (bool, error) {
	return false, nil
}

func logDBError(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("Error in [%s:%d] [%s]\n", file, line, err.Error())
}

func clearTransition(tx *sql.Tx) {
	rollbackRet := tx.Rollback()
	if rollbackRet != sql.ErrTxDone && rollbackRet != nil {
		util.LogError.Printf("Critical db error in rollback: %s", rollbackRet.Error())
	}
}
