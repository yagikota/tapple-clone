package infra

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/configs"
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

	// 確認用
	log.Println(dsn)

	return &MySQLConnector{
		Conn: conn,
	}
}

func mysqlConnInfo(mysqlInfo configs.MysqlInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassWord,
		mysqlInfo.MySQLAddr,
		mysqlInfo.MySQLDBName)

	return dataSourceName
}

type User struct {
	ID        int
	Name      string
	Icon      string
	CreatedAt time.Time
	DeletedAt sql.NullTime
}
