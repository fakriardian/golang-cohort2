package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"column:id; primaryKey" json:"id,omitempty"`
	CustomerName string    `gorm:"column:customer_name; not null" json:"customerName"`
	OrederedAt   time.Time `gorm:"column:ordered_at; not null" json:"orderedAt"`
	Items        []Item    `json:"items"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (Order) TableName() string {
	return "orders"
}
