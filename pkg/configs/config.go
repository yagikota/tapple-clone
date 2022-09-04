package configs

import (
	"os"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MysqlInfo
	AWSInfo   *AWSInfo
}

type HTTPInfo struct {
	Addr string
}

type MysqlInfo struct {
	MySQLUser     string
	MySQLPassWord string
	MySQLAddr     string
	MySQLDBName   string
}

type AWSInfo struct {
	S3 *S3
}

type S3 struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
}

func LoadConfig() *appConfig {
	addr := ":" + os.Getenv("PORT")
	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	dbInfo := &MysqlInfo{
		MySQLUser:     mysqlUser,
		MySQLPassWord: mysqlPassword,
		MySQLAddr:     mysqlAddr,
		MySQLDBName:   mysqlDBName,
	}

	S3 := &S3{
		Region:          "",
		Bucket:          "",
		AccessKeyID:     "",
		SecretAccessKey: "",
	}
	awsInfo := &AWSInfo{
		S3: S3,
	}

	conf := appConfig{
		MySQLInfo: dbInfo,
		HTTPInfo:  httpInfo,
		AWSInfo: awsInfo,
	}

	return &conf
}
