# Gorm-Metakit
Gorm-Metakit is a Go package designed to simplify pagination and sorting functionalities for applications using GORM (Go Object-Relational Mapping). This package provides a Metadata structure to manage pagination and sorting parameters, and a GORM scope function to apply these settings to database queries.

## Features
- Pagination: Easily handle pagination with customizable page size.
- Sorting: Sort query results by specifying sort parameters and direction.
- Default Settings: Provides sensible defaults for page, page size, and sort direction.


## Installation
To install the package, use go get:
```shell
go get github.com/Nicolas-ggd/gorm-metakit
```

## Usage
### Metadata Structure
The Metadata structure holds the pagination and sorting parameters:

```go
type Metadata struct {
    Page          int    `form:"page" json:"page"`
    PageSize      int    `form:"page_size" json:"page_size"`
    Sort          string `form:"sort" json:"sort"`
    SortDirection string `form:"sort_direction" json:"sort_direction"`
    TotalRows     int64  `json:"total_rows"`
    TotalPages    int64  `json:"total_pages"`
}
```

The simple way to bind Metadata struct is to use Gin for example.

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Nicolas-ggd/gorm-metakit"
)

func main() {
	r := gin.Default()

	r.GET("/items", func(c *gin.Context) {
		var metadata metakit.Metadata
		// Bind query parameters to the metadata struct
		if err := c.ShouldBind(&metadata); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	// Run the Gin server
	r.Run()
}
```

## Example Usage of Paginate function
```go
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fmt"
    "github.com/Nicolas-ggd/gorm-metakit"
)

func main() {
    // Initialize GORM DB connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Let's say that you already bind metakit.Metadata
    metadata := metakit.Metadata{
        Page:     1,
        PageSize: 20,
        Sort:     "name",
        SortDirection: "asc",
    }

    // Count total rows
    var totalRows int64
    db.Model(&Users{}).Count(&totalRows)
    metadata.TotalRows = totalRows

    // Fetch paginated and sorted results
    var results []Users
    db.Scopes(metakit.Paginate(&metadata)).Find(&results)

    fmt.Println(results)
}

```

## Contributing
This project is licensed under the MIT License. See the LICENSE file for details.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
