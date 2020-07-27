package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

const jsonPath = "mongo/json/"

func main() {
	filePaths := []string{jsonPath + "1000.json", jsonPath + "2100.json", jsonPath + "5050.json", jsonPath + "5500.json", jsonPath + "7000.json"}

	var dataSlice []mongo.Data
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
	start := time.Now()
	mongo.BulkInsert(dataSlice)
	end := time.Now()
	fmt.Printf("%f sec\n", (end.Sub(start)).Seconds())
}
