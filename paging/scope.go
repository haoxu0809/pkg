package paging

import (
	"github.com/haoxu0809/pkg/meta"

	"gorm.io/gorm"
)

func Paginate(opts *meta.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if opts.Page <= 0 {
			opts.Page = 1
		}
		switch {
		case opts.PageSize <= 0:
			opts.PageSize = 20
		case opts.PageSize > 100:
			opts.PageSize = 100
		}
		return db.Offset((opts.Page - 1) * opts.PageSize).Limit(opts.PageSize)
	}
}
