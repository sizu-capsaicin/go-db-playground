package main

import (
	"database/sql"
	"flag"
	"log"
)

const (
	mysqlPath = "tcp(127.0.0.1:3333)/"
	userName  = "root"
	passwd    = "xxxx"
)

func main() {
	var dbName = flag.String("db", "adb", "Database name")
	flag.Parse()

	path := userName + ":" + passwd + "@" + mysqlPath

	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatalln("Faild to DB connection.")
	}
	defer db.Close()

	// データベースの作成
	query := "create database " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Faild to exec query: %s\n", query)
	}

	// テーブルの作成
	query = "create table train()" + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Faild to exec query: %s\n", query)
	}
}
