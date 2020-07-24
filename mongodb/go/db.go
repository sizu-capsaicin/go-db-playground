package mongo

import (
	"encoding/json"
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	mongoPath = "mongodb://localhost/"
	dbName    = "adb"
	path      = mongoPath + dbName
)

// Insert は document を instert する関数
func Insert() {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	// json.RawMessage の変換
	for _, vnet := range vnets.VNet.VNetList {
		// 一旦 json に marshal
		op, err := json.Marshal(vnet.OverlayProperty)
		if err != nil {
			log.Fatalln(err)
		}

		err = bson.UnmarshalJSON(op, &vnet.OverlayProperty)
		if err != nil {
			log.Fatalln(err)
		}

		for _, vnode := range vnet.VNode {
			// 一旦 json に marshal
			rp, err := json.Marshal(vnode.ResourceProperty)
			if err != nil {
				log.Fatalln(err)
			}

			err = bson.UnmarshalJSON(rp, &vnode.ResourceProperty)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	col := db.C("vnet")
	if err := col.Insert(vnets); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
