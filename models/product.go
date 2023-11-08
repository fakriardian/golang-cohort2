package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID        uuid.UUID  `gorm:"column:id;type:varchar(191)" json:"id,omitempty"`
	Name      string     `gorm:"column:name; not null" json:"name"`
	ImageUrl  string     `gorm:"column:image_url; not null" json:"imageUrl"`
	UserID    uuid.UUID  `gorm:"type:varchar(191);" json:"-"`
	User      User       `json:"admin"`
	Variants  []Variant  `gorm:"foreignKey:ProductID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"variants"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}

func (Product) TableName() string {
	return "products"
}
