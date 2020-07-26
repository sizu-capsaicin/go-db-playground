package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlPath = "tcp(127.0.0.1:3306)/"
	userName  = "go"
)

func main() {
	var (
		dbName = flag.String("db", "adb", "Database name")
		passwd = flag.String("pw", "xxxx", "Enter the password")
	)
	flag.Parse()

	path := userName + ":" + *passwd + "@" + mysqlPath

	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// データベースのスイッチ
	query := "use " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// テーブルが既に作成済みであれば drop する
	query = "drop temporary table if exists `companies`, `safeties`, `lines`, `trains`"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// データベースが既に作成済みであれば drop する
	query = "drop database if exists " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
