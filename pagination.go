package metakit

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
func (m *Metadata) setPage() {
	if m.Page == 0 {
		m.Page = 1
	}
}

// SetPageSize function handle PageSize, first it's set default value 10. If page size is greater than 100, then it sets 100
func (m *Metadata) setPageSize() {
	switch {
	case m.PageSize > 100:
		m.PageSize = 100
	case m.PageSize <= 0:
		m.PageSize = 10
	}
}
