package main

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	// 読み出し用のスライスを定義
	var results []mongo.Data
	// 検索クエリ
	query := bson.M{"train.lines.railwidth": 1435}

	// DB からの読み出し
	mongo.FindAll(&results, query)
	var name map[string]string

	// 結果の表示
	var str string
	for _, r := range results {
		for _, l := range r.Train.Lines {
			if _, ok := name[l.Name]; ok == false {
				name[l.Name] = l.Name
				str += "  |- name: " + l.Name + "\n"
			}
		}
	}
	fmt.Print(str)
}
