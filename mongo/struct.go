package mongo

// Data は以下に続くデータを格納する構造体
type Data struct {
	Train Train `json:"train"`
}

// Train は train table のデータを格納する構造体
type Train struct {
	// 形式名
	Name string `json:"name"`
	// 運行路線
	Lines []Line `json:"lines"`
	// 運行会社
	Operator Company `json:"operator"`
	// 最高運転速度
	MaxSpeed int `json:"max-speed"`
	// 起動加速度
	Acceleration int `json:"acceleration"`
}

// Line は line table のデータを格納する構造体
type Line struct {
	// 路線名
	Name string `json:"name"`
	// 軌間
	RailWidth int `json:"rail-width"`
	// 路線距離
	RailRange int `json:"rail-range"`
	// 駅数
	Stations int `json:"stations"`
	// 保安装置
	Safeties []Safety `json:"safeties"`
	// 運営者
	Operator Company `json:"operator"`
}

// Company は company table のデータを格納する構造体
type Company struct {
	// 社名
	Name string `json:"name"`
	// 所在地
	Location string `json:"location"`
	// 公式サイト
	URL string `json:"URL"`
}

// Safety は safety table のデータを格納する構造体
type Safety struct {
	// 保安装置名
	Name string `json:"name"`
}
