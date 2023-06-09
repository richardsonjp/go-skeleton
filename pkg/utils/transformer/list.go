package transformer

type Pagination struct {
	Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Pagination PaginationMeta `json:"pagination"`
	MetaData   interface{}    `json:"meta_data,omitempty"`
}

type PaginationMeta struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}
