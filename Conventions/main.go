package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Base Model's definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type User struct {
	gorm.Model
	Name string
}

// Only need field `ID`, `CreatedAt`
type CustomUser struct {
	ID        uint
	CreatedAt time.Time
	Name      string
}
