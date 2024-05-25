package metakit

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math"
	"testing"
)

func TestSortDirectionParams(t *testing.T) {
	tests := []struct {
		input    Metadata
		expected string
	}{
		{Metadata{SortDirection: ""}, "asc"},
		{Metadata{SortDirection: "desc"}, "desc"},
	}

	for _, test := range tests {
		test.input.SortDirectionParams()
		if test.input.SortDirection != test.expected {
			t.Errorf("expected %v, got %v", test.expected, test.input.SortDirection)
		}
	}
}

func TestSortParams(t *testing.T) {
	m := Metadata{}
	sort := "name"
	m.SortParams(sort)
	if m.Sort != sort {
		t.Errorf("expected %v, got %v", sort, m.Sort)
	}
}

func TestPaginate(t *testing.T) {
	tests := []struct {
		metadata     Metadata
		expectedPage int
		expectedSize int
		expectedOff  int
	}{
		{Metadata{Page: 0, PageSize: 0, TotalRows: 100}, 1, 10, 0},
		{Metadata{Page: 3, PageSize: 20, TotalRows: 100}, 3, 20, 40},
		{Metadata{Page: 2, PageSize: 50, TotalRows: 120}, 2, 50, 50},
	}

	for _, test := range tests {
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			t.Fatalf("failed to connect to database: %v", err)
		}

		paginate := Paginate(&test.metadata)
		db = paginate(db)

		test.metadata.setPage()
		test.metadata.setPageSize()

		totalPages := int(math.Ceil(float64(test.metadata.TotalRows) / float64(test.metadata.PageSize)))
		if test.metadata.TotalPages != int64(totalPages) {
			t.Errorf("expected %v total pages, got %v", totalPages, test.metadata.TotalPages)
		}

		offset := (test.metadata.Page - 1) * test.metadata.PageSize
		if offset != test.expectedOff {
			t.Errorf("expected offset %v, got %v", test.expectedOff, offset)
		}
	}
}
