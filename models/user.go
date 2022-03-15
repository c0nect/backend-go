package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	ImageUrl  string    `json:"image_url"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt 	gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
