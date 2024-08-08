package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string
}
