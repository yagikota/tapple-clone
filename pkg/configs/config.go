package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	MySQLInfo *MysqlInfo
}

type MysqlInfo struct {
	Addr          string
	MySQLUser     string
	MySQLPassWord string
	MySQLAddr     string
	MySQLBDName   string
}

func init() {
	// ローカルでのみ必要。DockerでGo Serverを立ち上げる際はコメントアウト
	// ここから
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// ここまで
}

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
		MySQLBDName:   mysqlDBName,
	}

	conf := appConfig{
		MySQLInfo: &dbInfo,
	}

	return &conf
}
