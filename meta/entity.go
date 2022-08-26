package meta

import "gorm.io/plugin/soft_delete"

type Entity struct {
	Id        uint                  `json:"id" gorm:"primarykey"`
	CreatedAt uint                  `json:"createdAt"`
	UpdatedAt uint                  `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}

type Paginate struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
