package request

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" binding:"required" form:"id"` // 主键ID
}

type GetByIds struct {
	ID string `json:"ids" binding:"required" form:"ids"` // 主键ID
}
