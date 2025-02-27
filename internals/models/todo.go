package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Description string    `gorm:"not null" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRequest struct {
	Description string `json:"description" form:"description"`
}

type TodoResponse struct {
	ID uint `json:"id"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})

}
