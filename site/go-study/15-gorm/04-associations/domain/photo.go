package domain

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Description string
	UserID      uint
}
