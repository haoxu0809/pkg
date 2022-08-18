package meta

import "gorm.io/plugin/soft_delete"

type Entity struct {
	Id        uint                  `json:"id" gorm:"primarykey"`
	CreatedAt uint                  `json:"created_at"`
	UpdatedAt uint                  `json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}

type Paginate struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}
