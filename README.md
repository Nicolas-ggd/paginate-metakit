# Gorm-Metakit
Gorm-Metakit is a Go package designed to simplify pagination and sorting functionalities for applications using GORM (Go Object-Relational Mapping). This package provides a Metadata structure to manage pagination and sorting parameters, and a GORM scope function to apply these settings to database queries.

## Overview
- Pagination: Easily handle pagination with customizable page size.
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
## Example Usage of Paginate function

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
		err := c.ShoulBind(&metadata)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
			return
		}
		
		// Count total rows
		var totalRows int64
		db.Model(&Users{}).Count(&totalRows)
		metadata.TotalRows = totalRows

		// Fetch paginated and sorted results
		var results []Users
		db.Scopes(metakit.Paginate(&metadata)).Find(&results)
		
	})
}

```

## Contributing
This project is licensed under the MIT License. See the LICENSE file for details.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
