package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 環境変数読み込み
	addr := ":" + os.Getenv("PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", mysqlUser, mysqlPassword, mysqlAddr, mysqlDBName)
	mysqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer mysqlDB.Close()

	// mysqlとの接続確認
	if err := mysqlDB.Ping(); err != nil {
		panic(err)
	}

	// hello world api
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path); err != nil {
			fmt.Println(err)
		}
	})

	// starting server
	fmt.Printf("server is running. addr:%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
