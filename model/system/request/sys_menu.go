package request

//列表
type ParamMenuList struct {
	Page  int64  `json:"page" binding:"numeric"`
	Limit int64  `json:"limit" binding:"required,numeric"`
	Sort  string `json:"sort"`
	Title string `json:"title"`
}
