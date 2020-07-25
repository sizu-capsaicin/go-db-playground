package mongo

// Train は train table のデータを格納する構造体
type Train struct {
	// 形式名
	name string
	// 運行路線
	lines []Line
	// 運行会社
	operator Company
	// 最高運転速度
	maxSpeed int
	// 起動加速度
	acceleration int
	// 搭載保安装置
	safeties []Safety
}

// Line は line table のデータを格納する構造体
type Line struct {
	// 路線名
	name string
	// 軌間
	railWidth int
	// 路線距離
	railRange int
	// 駅数
	stations int
	// 保安装置
	safeties []Safety
	// 運営者
	operator Company
}

// Company は company table のデータを格納する構造体
type Company struct {
	// 社名
	name string
	// 所在地
	location string
	// 公式サイト
	url string
}

// Safety は safety table のデータを格納する構造体
type Safety struct {
	// 保安装置名
	name string
}
