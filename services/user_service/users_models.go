package user_service

import "time"

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
type address struct {
	Village  string `json:"village"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type UserLoginResponse struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
	Token    string `json:"token"`
}

type UserResponse struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	//Address  address `json:"address"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	FullName string `json:"full_name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Phone    int    `json:"phone" `
	Password string `json:"password" validate:"required,min=8,max=32"`
	CreateAt time.Time
	UpdateAt time.Time
}
type UserCreateResponse struct {
	Name     string `json:"name" `
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	//CreateAt time.Time `json:"createAt"`
}

type UserUpdateResponse struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}

type DeleteRequest struct {
	Email string `json:"email"`
}
