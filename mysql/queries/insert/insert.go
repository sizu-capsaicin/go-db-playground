package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlPath = "tcp(127.0.0.1:3306)/"
	userName  = "root"
	passwd    = "xxxx"
	dbName    = "adb"
	path      = userName + ":" + passwd + "@" + mysqlPath + dbName
)

func main() {
	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("DB connection error.")
	}
	defer db.Close()

	// クエリの実行
	query := ""
	ins, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	ins.Exec("")
}
