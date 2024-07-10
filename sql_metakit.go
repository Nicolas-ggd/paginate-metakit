package metakit

import (
	"database/sql"
	"fmt"
)

// SPaginate calculates the total pages and offset based on the current metadata and applies pagination to the SQL query
func SPaginate(db *sql.DB, query string, m *Metadata) (*sql.Rows, error) {
	m.setPage()
	m.setPageSize()

	// calculate the total pages
	if m.PageSize > 0 {
		totalPages := (m.TotalRows + int64(m.PageSize) - 1) / int64(m.PageSize)
		m.TotalPages = totalPages
	} else {
		m.TotalPages = 1
	}

	// calculate offset for the current page
	offset := (m.Page - 1) * m.PageSize

	// build the paginated query
	paginatedQuery := fmt.Sprintf("%s ORDER BY %s %s LIMIT %d OFFSET %d", query, m.Sort, m.SortDirection, m.PageSize, offset)
	
	rows, err := db.Query(paginatedQuery)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
