package main

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/sizu-capsaicin/go-db-playground/mongo"
	"golang.org/x/sync/errgroup"
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
	// サンプルデータの取得
	var r []mongo.Data
	q := bson.M{"train.name": "京急 2100 形"}

	// find
	mongo.FindAll(&r, q)

	// 一旦寝かせる
	time.Sleep(time.Second * 10)

	// max-speed の抜き出し
	maxSpeed := r[0].Train.MaxSpeed
	// セレクタ・更新内容の設定
	s := bson.M{"train.name": "京急 2100 形"}
	u := bson.M{"$set": bson.M{"train.maxspeed": maxSpeed + update}}

	// 一旦寝かせる
	time.Sleep(time.Second * 10)

	// 更新
	mongo.Update(s, u)
}
