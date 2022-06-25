package request

//列表
type ParamLoginLogList struct {
	Page     int64  `json:"page" binding:"numeric"`
	Limit    int64  `json:"limit" binding:"required,numeric"`
	UserName string `json:"username"`
	Status   string `json:"status"`
	Ip       string `json:"ip"`
}
