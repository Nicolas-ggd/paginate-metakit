package metakit

import (
	"gorm.io/gorm"
	"math"
)

type Metadata struct {
	// Page represents current page
	Page int `form:"page" json:"page"`

	// PageSize is capacity of per page items
	PageSize int `form:"page_size" json:"page_size"`

	// Sort is string type which defines the sort type of data
	Sort string `form:"sort" json:"sort"`

	// SortDirection defines sorted column name
	SortDirection string `form:"sort_direction" json:"sort_direction"`

	// TotalRows defines the quantity of total rows
	TotalRows int64 `json:"total_rows"`

	// TotalPages defines the quantity of total pages, it's defined based on page size and total rows
	TotalPages int64 `json:"total_pages"`
}

// SortDirectionParams function check SortDirection parameter, if it's empty, then it sets ascending order by default
func (m *Metadata) SortDirectionParams() {
	if m.SortDirection == "" {
		m.SortDirection = "asc"
	}
}

// SortParams function take string parameter of sort and set of Sort value
func (m *Metadata) SortParams(sort string) {
	m.Sort = sort
}

// SetPage function sets Page value as a 1 by default, if its equals to 0
func (m *Metadata) SetPage() {
	if m.Page == 0 {
		m.Page = 1
	}
}

// SetPageSize function handle PageSize, first it's set default value 10. If page size is greater than 100, then it sets 100
func (m *Metadata) SetPageSize() {
	switch {
	case m.PageSize > 100:
		m.PageSize = 100
	case m.PageSize <= 0:
		m.PageSize = 10
	}
}

// Paginate is GORM scope function. Paginate calculates the total pages and offset based on current metadata and applies pagination to the Gorm query
func Paginate(m *Metadata) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		m.SetPage()
		m.SetPage()

		// Calculate total pages based on total rows and page size
		totalPages := int(math.Ceil(float64(m.TotalRows) / float64(m.PageSize)))
		m.TotalPages = int64(totalPages)

		// Calculate offset for the current page
		offset := (m.Page - 1) * m.PageSize

		// Apply offset and limit to the Gorm query
		return db.Offset(offset).Limit(m.PageSize)
	}
}
