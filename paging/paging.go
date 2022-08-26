package paging

import (
	"github.com/haoxu0809/pkg/meta"

	"gorm.io/gorm"
)

func Transform(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize <= 0:
		pageSize = 20
	case pageSize > 100:
		pageSize = 100
	}

	return page, pageSize
}

func Paginate(opts *meta.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, pageSize := Transform(opts.Page, opts.PageSize)
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
