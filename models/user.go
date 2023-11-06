package models

import (
	"errors"
	"final-challenge/helpers"

	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `gorm:"column:id; primaryKey" json:"id,omitempty"`
	Name      string     `gorm:"column:name; not null" json:"name"`
	Email     string     `gorm:"column:email; unique; not null" json:"email"`
	Password  string     `gorm:"column:password; not null" json:"-"`
	Products  []Product  `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	encryption := helpers.NewEncryption()
	encrypted := encryption.Encrypt(u.Email)

	u.Email = encrypted
	u.Password, err = helpers.GenerateHashPassword(u.Password)
	if err != nil {
		return errors.New("invalid hash password")
	}

	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Email != "" {
		encryption := helpers.NewEncryption()
		decrypted := encryption.Decrypt(u.Email)

		u.Email = decrypted
	}

	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	encryption := helpers.NewEncryption()
	decrypted := encryption.Decrypt(u.Email)

	u.Email = decrypted
	return
}

func (User) TableName() string {
	return "users"
}
