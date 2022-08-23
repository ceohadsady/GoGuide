package user_repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	FullName string
	Email    string `gorm:"uniqueIndex"`
	Phone    int
	Password string
	//Address  address
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}
type address struct {
	Village  string
	City     string
	Province string
}
