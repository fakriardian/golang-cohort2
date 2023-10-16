package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Item struct {
	ID          uuid.UUID `gorm:"column:id; primaryKey" json:"id,omitempty"`
	Name        string    `gorm:"column:name; not null" json:"name"`
	Description string    `gorm:"column:description; not null" json:"description"`
	Quantity    int       `gorm:"column:quantity; not null" json:"quantity"`
	OrderID     uuid.UUID `gorm:"type:varchar(191); primaryKey" json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (Item) TableName() string {
	return "items"
}
