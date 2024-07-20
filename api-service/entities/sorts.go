package entities

type SortData struct {
	SortType string `json:sortType`
	DataList []int  `json:dataList`
}
