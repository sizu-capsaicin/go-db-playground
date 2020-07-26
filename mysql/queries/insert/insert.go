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
	query := "insert into `safeties`(name) values(?)"
	ins, err := db.Prepare(query)
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
	ins.Exec("1 号型 ATS")

	// companies への insert
	query = "insert into `companies`(name, location, url) values(?, ?, ?)"
	ins, err = db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	ins.Exec("東急電鉄", "東京都渋谷区神泉町 8 番 16 号", "https://www.tokyu.co.jp/railway/")
	ins.Exec("京浜急行電鉄", "神奈川県横浜市西区高島 1 丁目 2 番 8 号", "https://www.keikyu.co.jp")
	ins.Exec("横浜高速鉄道", "神奈川県横浜市中区元町 1 丁目 11 番", "http://www.mm21railway.co.jp/")
	ins.Exec("東京地下鉄", "東京都台東区東上野 3 丁目 19 番 6 号", "http://www.toukyometro.jp")
	ins.Exec("西武鉄道", "埼玉県所沢市くすのき台 1 丁目 11 番地の 1", "http://www.seiburailway.jp")
	ins.Exec("東武鉄道", "東京都墨田区押上 2 丁目 18 番 12 号", "http://www.tobu.co.jp")
	ins.Exec("京成本線", "千葉県市川市八幡 3 丁目 3 番 1 号", "https://www.keisei.co.jp")
	ins.Exec("北総鉄道", "千葉県鎌ケ谷市新鎌ヶ谷 4 丁目 2 番 3 号", "https://www.hokuso-railway.co.jp")
	ins.Exec("芝山鉄道", "千葉県山武郡芝山町香山新田 148 番地 1", "https://www.sibatetu.co.jp")
	ins.Exec("東京都交通局", "東京都新宿区西新宿 2 丁目 8 番 1 号", "https://www.kotsu.metro.tokyo.jp")

	// lines への insert
	query = "insert into `lines`(name, rail_width, rail_range, stations, safety_id, operator_id) values(?, ?, ?, ?, ?, ?)"
	ins, err = db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	ins.Exec("東急多摩川線", 1067, 5.6, 7, 1, 1)
	ins.Exec("東急池上線", 1067, 10.9, 15, 1, 1)
	ins.Exec("京急本線", 1435, 56.7, 50, 2, 2)
	ins.Exec("京急空港線", 1435, 6.5, 7, 2, 2)
	ins.Exec("京急久里浜線", 1435, 13.4, 9, 2, 2)
	ins.Exec("京急逗子線", 1435, 5.9, 4, 2, 2)
	ins.Exec("東急東横線", 1067, 24.2, 21, 3, 1)
	ins.Exec("横浜高速鉄道みなとみらい線", 1067, 4.1, 6, 3, 3)
	ins.Exec("東京メトロ副都心線", 1067, 11.9, 11, 4, 4)
	ins.Exec("東京メトロ副都心線", 1067, 11.9, 11, 5, 4)
	ins.Exec("西武有楽町線", 1067, 2.6, 3, 6, 5)
	ins.Exec("西武池袋線", 1067, 57.8, 31, 7, 5)
	ins.Exec("東武東上本線", 1067, 75.0, 38, 8, 6)
	ins.Exec("東武東上本線", 1067, 75.0, 38, 9, 6)
	ins.Exec("京成押上線", 1435, 5.7, 6, 2, 7)
	ins.Exec("京成本線", 1435, 69.3, 42, 2, 7)
	ins.Exec("京成成田空港線 (成田スカイアクセス線)", 1435, 51.4, 8, 2, 7)
	ins.Exec("京成東成田線", 1435, 7.1, 2, 2, 7)
	ins.Exec("北総鉄道北総線", 1435, 32.3, 15, 2, 8)
	ins.Exec("芝山鉄道線", 1435, 2.2, 2, 2, 8)

}
