package configs

import (
	"os"
)

type appConfig struct {
	MySQLInfo *MysqlInfo
}

type MysqlInfo struct {
	Addr          string
	MySQLUser     string
	MySQLPassWord string
	MySQLAddr     string
	MySQLDBName   string
}

// TODO: 必要か調べる？
// func init() {
// }

func LoadConfig() *appConfig {
	addr := ":" + os.Getenv("PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dbInfo := MysqlInfo{
		Addr:          addr,
		MySQLUser:     mysqlUser,
		MySQLPassWord: mysqlPassword,
		MySQLAddr:     mysqlAddr,
		MySQLDBName:   mysqlDBName,
	}

	conf := appConfig{
		MySQLInfo: &dbInfo,
	}

	return &conf
}
