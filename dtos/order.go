package dtos

type Item struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type OrderRequest struct {
	OrderedAt    string `json:"orderedAt" binding:"required"`
	CustomerName string `json:"customerName" binding:"required"`
	Items        []Item `json:"items" binding:"dive"`
}
