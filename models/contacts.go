package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Contact struct {
		ID        uuid.UUID `gorm:"primary_key" example:"2c45e4ec-26e0-4043-86e4-c15b9cf985a2" json:"id" binding:"max=63"`
		Name      string    `json:"name" binding:"required"`
		Gender    string    `json:"gender"`
		Phone     string    `json:"phone" binding:"max=12"`
		Email     string    `json:"email" binding:"required"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
