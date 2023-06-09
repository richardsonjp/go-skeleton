package model

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Limit > 200 {
		p.Limit = 200
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "updated_at desc"
	}
	return p.Sort
}

func (p *Pagination) GetTotalPages() int {
	return int(math.Ceil(float64(p.TotalRows) / float64(p.Limit)))
}

func NewPaginate(limit int, page int) *Pagination {
	return &Pagination{Limit: limit, Page: page}
}

func (p *Pagination) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit

	return db.Offset(offset).
		Limit(p.Limit)
}
