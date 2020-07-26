package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	// セレクタ・更新内容の設定
	s := bson.M{"train.name": "京急 2100 形"}
	u := bson.M{"$set": bson.M{"train.maxspeed": 110}}

	// 更新
	mongo.Update(s, u)
}
