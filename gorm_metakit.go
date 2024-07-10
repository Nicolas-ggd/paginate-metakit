package metakit

import (
	"gorm.io/gorm"
)

// GPaginate is GORM scope function. Paginate calculates the total pages and offset based on current metadata and applies pagination to the Gorm query
// GPaginate function cares Page and PageSize automatically, you can use your own function to replace it, it just overwrite fields
func GPaginate(m *Metadata) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		m.setPage()
		m.setPageSize()

		// Use integer arithmetic to avoid possible unsafe checking
		if m.PageSize > 0 {
			totalPages := (m.TotalRows + int64(m.PageSize) - 1) / int64(m.PageSize)
			m.TotalPages = totalPages
		} else {
			m.TotalPages = 1
		}

		// Calculate offset for the current page
		offset := (m.Page - 1) * m.PageSize

		// Apply offset and limit to the Gorm query
		return db.Offset(offset).Limit(m.PageSize)
	}
}
