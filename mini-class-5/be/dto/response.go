package dto

type Response struct {
	Data    []interface{} `json:"data"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
}
