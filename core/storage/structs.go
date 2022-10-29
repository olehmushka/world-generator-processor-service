package storage

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type SortField struct {
	FieldName string `json:"field_name"`
	IsDESC    bool   `json:"is_desc"`
}

type PaginationSortingOpts struct {
	Sorting    []*SortField `json:"sorting"`
	Pagination *Pagination  `json:"pagination"`
}
