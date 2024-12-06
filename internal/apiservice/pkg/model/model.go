package model

import "time"

type Model struct {
	ID        int64     `json:"id" gorm:"primaryKey" `
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
