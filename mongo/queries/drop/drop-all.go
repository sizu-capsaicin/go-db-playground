package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	// DB の Document を Drop
	mongo.DropAll(bson.M{})
}
