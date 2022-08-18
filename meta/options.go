package meta

type ListOptions struct {
	Page     int    `form:"page" binding:"omitempty"`
	PageSize int    `form:"page_size" binding:"omitempty"`
	Order    string `form:"order" binding:"omitempty,oneof=asc desc"`
	OrderBy  string `form:"order_by" binding:"omitempty,oneof=created_at updated_at"`
}
