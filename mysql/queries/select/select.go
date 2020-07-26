package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	mysql "github.com/sizu-capsaicin/go-db-playground/mysql"
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
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}

	var results []mysql.Train
	for rows.Next() {
		var train mysql.Train
		if err = rows.Scan(train.Name); err != nil {
			log.Fatalln(err)
		}
		results = append(results, train)
	}

	for _, r := range results {
		fmt.Printf("Train name: %s\n", r.Name)
	}
}
