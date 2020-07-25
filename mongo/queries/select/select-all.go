package main

import (
	"fmt"
	"strconv"

	"github.com/sizu-capsaicin/go-db-playground/mongo"
)

func main() {
	// 読み出し用のスライスを定義
	var results []mongo.Data

	// DB からの読み出し
	mongo.FindAll(&results)

	// 結果の表示
	var str string
	for _, r := range results {
		t := r.Train
		str += "Train\n"
		str += "|- name: " + t.Name + "\n"
		str += "|- lines\n"
		for _, l := range t.Lines {
			str += "  |- name: " + l.Name + "\n"
			str += "  |- raild-width: " + strconv.Itoa(l.RailWidth) + "\n"
			str += "  |- raild-range: " + strconv.FormatFloat(l.RailRange, 'f', 1, 64) + "\n"
			str += "  |- stations: " + strconv.Itoa(l.Stations) + "\n"
			str += "  |- safeties\n"
			for _, s := range l.Safeties {
				str += "    |- name: " + s.Name + "\n"
			}
			str += "  |- operator\n"
			str += "    |- name: " + l.Operator.Name + "\n"
			str += "    |- location: " + l.Operator.Location + "\n"
			str += "    |- URL: " + l.Operator.URL + "\n"
		}
		str += "|- operator\n"
		str += "  |- name: " + t.Operator.Name + "\n"
		str += "  |- location: " + t.Operator.Location + "\n"
		str += "  |- URL: " + t.Operator.URL + "\n"
		str += "|- max-speed: " + strconv.Itoa(t.MaxSpeed) + "\n"
		str += "|- acceleration: " + strconv.FormatFloat(t.Acceleration, 'f', 1, 64) + "\n"
	}
	fmt.Print(str)
}
