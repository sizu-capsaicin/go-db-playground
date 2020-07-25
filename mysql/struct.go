package mysql

// Train は train table のデータを格納する構造体
type Train struct {
	// 形式名
	Name string
	// 運行路線
	Lines []Line
	// 運行会社
	Operator Company
	// 最高運転速度
	MaxSpeed int
	// 起動加速度
	Acceleration int
	// 搭載保安装置
	Safeties []Safety
}

// Line は line table のデータを格納する構造体
type Line struct {
	// 路線名
	Name string
	// 軌間
	RailWidth int
	// 路線距離
	RailRange int
	// 駅数
	Stations int
	// 保安装置
	Safeties []Safety
	// 運営者
	Operator Company
}

// Company は company table のデータを格納する構造体
type Company struct {
	// 社名
	Name string
	// 所在地
	Location string
	// 公式サイト
	URL string
}

// Safety は safety table のデータを格納する構造体
type Safety struct {
	// 保安装置名
	Name string
}
