package main

import (
	"fmt"
	"net/http"

	"github.com/CyberAgentHack/2208-ace-go-server/configs"
	"github.com/CyberAgentHack/2208-ace-go-server/infra"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	_ = infra.NewMySQLConnector()

	// hello world api 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Hello, C Team. you've requested: %s\n", r.URL.Path); err != nil {
			fmt.Println(err)
		}
	})

	// hello world api 2
	http.HandleFunc("/hoge", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Hello, C Team. you've requested: %s\n", r.URL.Path); err != nil {
			fmt.Println(err)
		}
	})

	addr := configs.LoadConfig().MySQLInfo.Addr
	// starting server
	fmt.Printf("server is running. addr:%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
