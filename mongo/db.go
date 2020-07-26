package mongo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	mongoPath = "mongodb://localhost/"
	dbName    = "adb"
	path      = mongoPath + dbName
	cName     = "train"
)

// Insert は document を instert する関数
func Insert(data Data) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	col := db.C(cName)
	if err := col.Insert(data); err != nil {
		log.Fatalln(err)
	}
}

// BulkInsert は document を bulk instert する関数
func BulkInsert(data []Data) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	col := db.C(cName)
	bulk := col.Bulk()
	for _, d := range data {
		bulk.Insert(d)
	}

	br, err := bulk.Run()
	if err != nil {
		log.Fatalln(err)
	}
	if br != nil {
		fmt.Printf("bulk result: %+v\n", br)
	}
}

// DropAll は指定した collection の全 document を delete する関数
func DropAll(query bson.M) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	info, err := db.C(cName).RemoveAll(query)
	if err != nil {
		log.Fatalln(err)
	}
	if info != nil {
		fmt.Printf("Removed: %d, Updated: %d\n", info.Removed, info.Updated)
	}
}

// FindAll は全ての document を select する関数
func FindAll(data interface{}, query bson.M) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	db.C(cName).Find(query).All(data)
}

// Update はセレクタに該当する document を全て update する関数
func Update(selector interface{}, update interface{}) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	// 更新
	info, err := db.C(cName).UpdateAll(selector, update)
	if err != nil {
		log.Fatalln(err)
	}
	if info != nil {
		fmt.Printf("%+v\n", info)
	}
}
