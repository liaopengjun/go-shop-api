package request

//列表
type ParamSysApiList struct {
	Page     int64  `json:"page" binding:"numeric"`
	Limit    int64  `json:"limit" binding:"required,numeric"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	ApiGroup string `json:"api_group"`
}

// 表单参数
type ParamSysApiData struct {
	Id          uint   `json:"id"`
	Path        string `json:"path" binding:"required"`
	Method      string `json:"method" binding:"required"`
	ApiGroup    string `json:"api_group" binding:"required"`
	Description string `json:"description"`
}
