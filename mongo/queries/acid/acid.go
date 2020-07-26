package main

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sizu-capsaicin/go-db-playground/mongo"
	"golang.org/x/sync/errgroup"
)

const (
	mongoPath = "mongodb://localhost/"
	dbName    = "adb"
	path      = mongoPath + dbName
	cName     = "train"
)

func main() {
	// 並行処理
	eg := errgroup.Group{}
	eg.Go(func() error {
		findAndUpdate(-10)
		return nil
	})
	eg.Go(func() error {
		findAndUpdate(+10)
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}
}

func findAndUpdate(update int) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	// サンプルデータの取得
	var r []mongo.Data
	q := bson.M{"train.name": "京急 2100 形"}

	// find
	db.C(cName).Find(q).All(&r)

	// max-speed の抜き出し
	maxSpeed := r[0].Train.MaxSpeed
	// セレクタ・更新内容の設定
	selector := bson.M{"train.name": "京急 2100 形"}
	u := bson.M{"$set": bson.M{"train.maxspeed": maxSpeed + update}}

	// 一旦寝かせる
	time.Sleep(time.Second * 10)

	// 更新
	info, err := db.C(cName).UpdateAll(selector, u)
	if err != nil {
		log.Fatalln(err)
	}
	if info != nil {
		fmt.Printf("%+v\n", info)
	}
}
