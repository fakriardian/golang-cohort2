package dtos

type Response struct {
	AccessToken string        `json:"accessToken,omitempty"`
	Data        []interface{} `json:"data,omitempty"`
	Error       string        `json:"error,omitempty"`
	Total       int64         `json:"total,omitempty"`
	Page        int64         `json:"page,omitempty"`
	PageSize    int64         `json:"pageSize,omitempty"`
	Status      int64         `json:"status"`
	Message     string        `json:"message"`
}
