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
	dbName    = "adb"
)

func main() {
	var passwd = flag.String("pw", "xxxx", "Enter the password")
	flag.Parse()

	// MySQL への接続
	path := userName + ":" + *passwd + "@" + mysqlPath + dbName
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("DB connection error.")
	}
	defer db.Close()

	// クエリの実行
	// safeties への insert
	q := "insert into `safeties`(name) values(?)"
	ins, err := db.Prepare(q)
	if err != nil {
		log.Fatalln(err)
	}
	ins.Exec("東急 ATS")
	ins.Exec("C-ATS")
	ins.Exec("ATC-P")
	ins.Exec("新 CS-ATS")
	ins.Exec("ATO")
	ins.Exec("CS-ATC (ATC-4型)")
	ins.Exec("ATS")
	ins.Exec("T-DATC")
	ins.Exec("東武 ATS (TSP)")
}
