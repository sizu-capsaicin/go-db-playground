package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

const jsonPath = "mongo/json/"

func main() {
	filePaths := []string{jsonPath + "1000.json", jsonPath + "2100.json", jsonPath + "5050.json", jsonPath + "5500.json", jsonPath + "7000.json"}

	dataSlice := make([]mongo.Data, 5)
	for _, fp := range filePaths {
		// json ファイルを開く
		file, err := ioutil.ReadFile(fp)
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
		dataSlice = append(dataSlice, d)
	}

	// DB への書き込み
	mongo.BulkInsert(dataSlice)
}
