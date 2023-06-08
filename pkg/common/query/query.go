package query

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// KeyWords ...
type KeyWords = map[string]interface{}

// Query parameters
type Query struct {
	// Filter list
	Keywords KeyWords
	// Sort list
	Sorts []*Sort
	// Page number
	PageNumber int
	// Page size
	PageSize int
	// Without pagination
	WithoutPagination bool
}

// Sort specifies the order information
type Sort struct {
	Key  string
	DESC bool
}

// NewSort creates new sort
func NewSort(key string, desc bool) *Sort {
	return &Sort{
		Key:  key,
		DESC: desc,
	}
}

// Range query
type Range struct {
	Min interface{}
	Max interface{}
}

// NewRange creates a new range
func NewRange(min, max interface{}) *Range {
	return &Range{
		Min: min,
		Max: max,
	}
}

func (q *Query) Limit() int {
	if q.WithoutPagination {
		return MaxItems
	}
	if q.PageSize < 1 {
		q.PageSize = DefaultPageSize
	}
	return q.PageSize
}

func (q *Query) Offset() int {
	if q.WithoutPagination {
		return 0
	}
	if q.PageSize < 1 {
		q.PageSize = DefaultPageSize
	}
	if q.PageNumber < 1 {
		q.PageNumber = DefaultPageNumber
	}
	return (q.PageNumber - 1) * q.PageSize
}

func (q *Query) WithPagination(c *gin.Context) *Query {
	pageNumberStr := c.Query(PageNumber)
	pageNumber, _ := strconv.Atoi(pageNumberStr)

	pageSizeStr := c.Query(PageSize)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if pageNumber < 1 {
		pageNumber = DefaultPageNumber
	}

	if pageSize < 0 {
		pageSize = DefaultPageSize
	}

	q.PageSize = pageSize
	q.PageNumber = pageNumber
	return q
}
