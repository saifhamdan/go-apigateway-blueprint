// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Query struct {
	QueryString string
	Page        int
	Limit       int
	SortBy      string
	IsEmpty     bool
}

// QueryFilter parses the query string and returns the query string for MySQL
// what it filters is the date range, page number and page limit (max 500)
func QueryFilter(c *fiber.Ctx) (*Query, error) {
	// empty query string at the start
	queryString := ""
	startsWithAnd := false

	// from date
	from := c.Query("from")
	if from != "" {
		fromTime, err := time.Parse(time.RFC3339, from)
		if err != nil {
			return nil, err
		}
		if startsWithAnd {
			queryString += " AND "
		}
		queryString += fmt.Sprintf("created_at >= %d", fromTime.UnixNano())
		startsWithAnd = true
	}

	// to date
	to := c.Query("to")
	if to != "" {
		toTime, err := time.Parse(time.RFC3339, to)
		if err != nil {
			return nil, err
		}

		if startsWithAnd {
			queryString += " AND "
		}
		queryString += fmt.Sprintf("created_at <= %d", toTime.UnixNano())
	}

	// page number
	page := c.QueryInt("page", 1)
	page = page - 1

	// set default limit value
	defaultLimitValue := 500

	// offset number
	limit := c.QueryInt("limit", defaultLimitValue)

	// if limit is more than 500, return an error and if it's 0 ignore it
	if limit > 0 && limit > defaultLimitValue {
		return nil, fmt.Errorf("you can't limit more than %d record per request", defaultLimitValue)
	}

	sortBy := c.Query("sort_by")

	return &Query{
		QueryString: queryString,
		Page:        page,
		Limit:       limit,
		SortBy:      sortBy,
		IsEmpty:     !startsWithAnd,
	}, nil
}
