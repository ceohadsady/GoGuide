package user_repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	FullName string `validate:"required,min=3,max=32"`
	Email    string `gorm:"uniqueIndex" `
	Phone    int
	Password string
	//Address
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}
type address struct {
	Village  string
	City     string
	Province string
}
