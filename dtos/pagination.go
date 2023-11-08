package dtos

import "strconv"

type PaginationResponse struct {
	Data     []interface{} `json:"data"`
	Error    string        `json:"error,omitempty"`
	Total    int64         `json:"total,omitempty"`
	Page     int64         `json:"page,omitempty"`
	PageSize int64         `json:"pageSize,omitempty"`
	Status   int64         `json:"status"`
	Message  string        `json:"message"`
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
