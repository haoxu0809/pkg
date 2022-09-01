package paging

import (
	"gorm.io/gorm"
)

func Transform(page, pageSize int64) (int64, int64) {
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

func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, pageSize := Transform(page, pageSize)
		return db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize))
	}
}
