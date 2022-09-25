package infra

import (
	"database/sql"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/configs"
	"github.com/go-sql-driver/mysql"
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

	// TODO: 取り扱い調べる
	// defer conn.Close()

	// mysqlとの接続確認
	if err := conn.Ping(); err != nil {
		panic(err)
	}

	return &MySQLConnector{
		Conn: conn,
	}
}

func mysqlConnInfo(mysqlInfo configs.MysqlInfo) string {
	cfg := mysql.Config{
		User:                 mysqlInfo.MySQLUser,
		Passwd:               mysqlInfo.MySQLPassWord,
		Addr:                 mysqlInfo.MySQLAddr,
		DBName:               mysqlInfo.MySQLDBName,
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}
	return cfg.FormatDSN()
}
