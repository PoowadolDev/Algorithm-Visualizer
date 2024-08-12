package entities

type SortData struct {
	SortType string `json:sortType`
	DataList []int  `json:dataList`
}

type SolveData struct {
	Step       string `json:step`
	DataSwitch []int  `json:dataList`
	DataList   []int  `json:dataList`
}
