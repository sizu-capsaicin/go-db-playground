package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	var json = flag.String("json", "mongo/json/2100.json", "json file path")

	// json ファイルを開く
	file, err := ioutil.ReadFile(json)
	if err != nil {
		panic(err)
	}

	// json を構造体へパース
	var t mongo.Train
	if err := json.Unmarshal(file, &t); err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			log.Fatalln(string(file[err.Offset-15 : err.Offset+15]))
		}
		log.Fatal(err)
	}

	// DB への書き込み
	mongo.Insert(t)
}
