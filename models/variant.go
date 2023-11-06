package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Variant struct {
	ID        uuid.UUID  `gorm:"column:id; primaryKey" json:"id,omitempty"`
	Name      string     `gorm:"column:name; not null" json:"name"`
	Quantity  int        `gorm:"column:quantity; not null" json:"quantity"`
	ProductID uuid.UUID  `gorm:"type:uuid; primarykey" json:"productId"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}

func (Variant) TableName() string {
	return "variants"
}
