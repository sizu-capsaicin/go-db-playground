package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Train は train table のデータを格納する構造体
type Train struct {
	number int
	lines  []Line
}

// Line は line table のデータを格納する構造体
type Line struct {
}

// Company は company table のデータを格納する構造体
type Company struct {
}

// Safety は safety table のデータを格納する構造体
type Safety struct {
}

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
	ins, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}

	ins.Exec("")
}
