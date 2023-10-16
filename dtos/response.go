package dtos

type Response struct {
	Data    []interface{} `json:"data,omitempty"`
	Error   string        `json:"error,omitempty"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
}
