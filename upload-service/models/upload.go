package models

import "time"

type Upload struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId    string    `gorm:"type:uuid;not null;references:users(id);constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	VideoKey  string    `gorm:"type:text;" json:"video_key"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}