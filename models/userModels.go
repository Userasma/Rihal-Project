package models

import (
	"time"

	"gorm.io/gorm"
)

type Main struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	DateOfBirth time.Time
	Age         int `gorm:"-"`
}
