# Pagination-Metakit
Pagination-Metakit is a Go package designed to simplify pagination and sorting functionalities for applications using GORM (Go Object-Relational Mapping) and Pure SQL. This package provides a Metadata structure to manage pagination and sorting parameters, and a GORM scope function to apply these settings to database queries, but not for only GORM, Pagination-Metakit also support pure sql pagination.

## Overview
- Pagination: Easily handle pagination with customizable page size.
- Default Settings: Provides sensible defaults for page, page size, and sort direction.
- Dual Functionality: Supports both GORM and pure SQL pagination.

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
## Example Usage of GPaginate function

```go
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"fmt"
	"github.com/Nicolas-ggd/gorm-metakit"
	"net/http"
)

func main() {
	r.GET("/items", func(c *gin.Context) {
		var metadata metakit.Metadata

		// Bind metakit Metadata struct 
		err := c.ShouldBind(&metadata)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
			return
		}

		// Count total rows
		var totalRows int64
		db.Model(&YourModel{}).Count(&totalRows)
		metadata.TotalRows = totalRows

		// Fetch paginated and sorted results
		var results []YourModel
		db.Scopes(metakit.GormPaginate(&metadata)).Find(&results)

		c.JSON(http.StatusOK, gin.H{"metadata": metadata, "results": results})
	})

	r.Run()
}

```

## Example usage of SQLPaginate function
```go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Nicolas-ggd/gorm-metakit"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    r := gin.Default()
    
    r.GET("/items", func(c *gin.Context) {
        var metadata metakit.Metadata
        
        // Bind metakit Metadata struct
        err := c.ShouldBind(&metadata)
        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
            return
        }
        
        // Count total rows
        row := db.QueryRow("SELECT COUNT(*) FROM your_table")
        err = row.Scan(&metadata.TotalRows)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            return
        }

        // Fetch paginated and sorted results
        query := "SELECT * FROM your_table"
        rows, err := metakit.SQLPaginate(db, query, &metadata)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            return
        }
        defer rows.Close()

        var results []YourModel
        for rows.Next() {
            var item YourModel
            // Scan your results here
            results = append(results, item)
        }
        
        c.JSON(http.StatusOK, gin.H{"metadata": metadata, "results": results})
    })

    r.Run()
}

```

## Contributing
This project is licensed under the MIT License. See the LICENSE file for details.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
