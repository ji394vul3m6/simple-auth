package util

const (
	mysqlSQLURLKey      = "AUTH_DB_URL"
	mysqlSQLPortKey     = "AUTH_DB_PORT"
	mysqlSQLUserKey     = "AUTH_DB_USER"
	mysqlSQLPasswordKey = "AUTH_DB_PASS"
	mysqlSQLDatabaseKey = "AUTH_DB"
)

// GetMySQLConfig will get db init config from env
func GetMySQLConfig() (url string, port int, user string, password string, database string) {
	url = GetStrEnv(mysqlSQLURLKey, "127.0.0.1")
	port = GetIntEnv(mysqlSQLPortKey, 3306)
	user = GetStrEnv(mysqlSQLUserKey, "root")
	password = GetStrEnv(mysqlSQLPasswordKey, "password")
	database = GetStrEnv(mysqlSQLDatabaseKey, "auth")
	return
}
