package models

import "gorm.io/gorm"

type Main struct {
	gorm.Model
	ClassID     int
	CountryID   int
	Name        string
	DateOfBirth string
}

type Class struct {
	gorm.Model
	ClassName string
}

type Student struct {
	gorm.Model
	Name string
}
