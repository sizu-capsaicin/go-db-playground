package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sizu-capsaicin/go-db-playground/mysql"
	"golang.org/x/sync/errgroup"
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

	// 並行処理
	eg := errgroup.Group{}
	eg.Go(func() error {
		findAndUpdate(-10, path)
		return nil
	})
	eg.Go(func() error {
		findAndUpdate(+10, path)
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}
}

func findAndUpdate(update int, path string) {
	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("DB connection error.")
	}
	defer db.Close()

	// transaction
	transaction, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	// find
	query := "select t.max_speed from `trains` as t where t.name=\"京急 2100 形\""
	rows, err := transaction.Query(query)
	if err != nil {
		log.Fatalln(err)
	}

	var results []mysql.Train
	for rows.Next() {
		var train mysql.Train
		if err = rows.Scan(&train.MaxSpeed); err != nil {
			log.Fatalln(err)
		}
		results = append(results, train)
	}
	maxSpeed := results[0].MaxSpeed

	// update
	query = "update trains set max_speed=? where name=?"
	upd, err := transaction.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	// sleep
	time.Sleep(time.Second * 10)

	r, err := upd.Exec(maxSpeed+update, "京急 2100 形")
	if err != nil {
		log.Fatalln(err)
	}
	if r != nil {
		fmt.Println(r)
	}

	transaction.Commit()
}
