package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	BirthDay time.Time
}
