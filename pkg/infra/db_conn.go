package infra

import (
	"database/sql"
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/configs"
)

const driverName = "mysql"

type MySQLConnector struct {
	Conn *sql.DB
}

// TODO: 返り値をinterfaceにする?
func NewMySQLConnector() *MySQLConnector {
	conf := configs.LoadConfig()

	dsn := mysqlConnInfo(*conf.MySQLInfo)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// mysqlとの接続確認
	if err := conn.Ping(); err != nil {
		panic(err)
	}

	return &MySQLConnector{
		Conn: conn,
	}
}

func mysqlConnInfo(mysqlInfo configs.MysqlInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassWord,
		mysqlInfo.MySQLAddr,
		mysqlInfo.MySQLBDName)

	return dataSourceName
}
