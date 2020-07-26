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
	q := "insert into `safeties`(name) values(?)"
	ins, err := db.Prepare(q)
	if err != nil {
		log.Fatalln(err)
	}

	ins.Exec("C-ATS")
}
