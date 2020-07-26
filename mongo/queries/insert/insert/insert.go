package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	var filePath = flag.String("json", "mongo/json/2100.json", "json file path")
	flag.Parse()

	// json ファイルを開く
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	// json を構造体へパース
	var d mongo.Data
	if err := json.Unmarshal(file, &d); err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			log.Fatalln(string(file[err.Offset-15 : err.Offset+15]))
		}
		log.Fatal(err)
	}

	// DB への書き込み
	mongo.Insert(d)
}
