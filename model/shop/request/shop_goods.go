package request

type GoodsParam struct {
	GoodsType       int    `json:"GoodsType"`
	PageNumber      int    `json:"pageNumber"`
	GoodsCategoryId string `json:"GoodsCategoryId"`
	Keyword         string `json:"Keyword"`
	OrderBy         string `json:"orderBy"`
}
