package main

import (
	"database/sql"
	"log"
)

const (
	mysqlPath = "tcp(127.0.0.1:3333)/"
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
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}

	var result []Train
	for _, row := range rows {
		train := Train{}

	}
}
