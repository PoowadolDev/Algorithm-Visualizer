package entities

type SortData struct {
	SortType string `json:sortType`
	DataList []int  `json:dataList`
}

type SolveData struct {
	Step     int   `json:step`
	DataList []int `json:dataList`
}
