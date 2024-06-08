package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FullName string    `gorm:"size:255;not null" json:"full_name"`
	Email    string    `gorm:"size:255;not null;unique" json:"email"`
	Password string    `gorm:"size:255; not null" json:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.New().String()
	return
}