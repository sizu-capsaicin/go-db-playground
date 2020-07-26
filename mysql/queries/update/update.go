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
	path := userName + ":" + *passwd + "@" + mysqlPath + dbName

	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("DB connection error.")
	}
	defer db.Close()

	// クエリの実行
	query := "update trains set max_speed=? where name=?"
	upd, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	upd.Exec(130, "京急 2100 形")
}
