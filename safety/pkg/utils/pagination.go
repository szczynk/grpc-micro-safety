package utils

import (
	"math"
	"strings"
)

const (
	defaultLimit uint32 = 5
	defaultSort         = "created_at desc"
)

// Pagination query params
type Pagination struct {
	Limit uint32 `json:"limit,omitempty"`
	Page  uint32 `json:"page,omitempty"`
	Sort  string `json:"sort,omitempty"`
}

// NewPaginationQuery Pagination query constructor
func NewPaginationQuery(limit uint32, page uint32, sort string) *Pagination {
	return &Pagination{Limit: limit, Page: page, Sort: sort}
}

// GetLimit Get limit
func (q *Pagination) GetLimit() uint32 {
	if q.Limit == 0 {
		q.Limit = defaultLimit
	}
	return q.Limit
}

// GetPage Get Page
func (q *Pagination) GetPage() uint32 {
	if q.Page == 0 {
		q.Page = 1
	}
	return q.Page
}

// GetSort Get Sort
func (q *Pagination) GetSort() string {
	if q.Sort == "" {
		q.Sort = defaultSort
	}
	return strings.Replace(q.Sort, ":", " ", -1)
}

// GetOffset Get offset
func (q *Pagination) GetOffset() uint32 {
	return (q.GetPage() - 1) * q.GetLimit()
}

// GetTotalPages Get total pages int
func (q *Pagination) GetTotalPages(totalCount uint32) uint32 {
	d := float64(totalCount) / float64(q.GetLimit())
	return uint32(math.Ceil(d))
}

// GetHasMore Get has more
func (q *Pagination) GetHasMore(totalCount uint32) bool {
	return q.GetPage() < q.GetTotalPages(totalCount)
}
