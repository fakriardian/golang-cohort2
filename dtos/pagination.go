package dtos

import "strconv"

type PaginationResponse struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Data     []interface{} `json:"data,omitempty"`
}

type FilterRequest struct {
	Search string `form:"search" binding:"-"`
	Page   string `form:"page" binding:"-"`
	Size   string `form:"size" binding:"-"`
}

type Filter struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Search string `json:"search"`
}

func Paginate(flr FilterRequest) Filter {
	var filter Filter
	if flr.Page == "" {
		filter.Page = 1
	} else {
		page, _ := strconv.Atoi(flr.Page)
		filter.Page = page
	}

	if flr.Size == "" {
		filter.Size = 10
	} else {
		size, _ := strconv.Atoi(flr.Size)
		if size > 30 {
			filter.Size = 30
		} else {
			filter.Size = size
		}
	}

	filter.Search = flr.Search

	return filter
}
