package user_service

import "time"

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	//Address  address `json:"address"`
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
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	//Address  address `json:"address"`
	CreateAt time.Time
	UpdateAt time.Time
}
type UserCreateResponse struct {
	Name     string  `json:"name"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Phone    int     `json:"phone"`
	Password string  `json:"password"`
	Address  address `json:"address"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}

type UserUpdateRequest struct {
	Name     string  `json:"name"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Phone    int     `json:"phone"`
	Password string  `json:"password"`
	Address  address `json:"address"`
	//CreateAt time.Time `json:"createAt"`
}

type UserUpdateResponse struct {
	Name     string  `json:"name"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Phone    int     `json:"phone"`
	Password string  `json:"password"`
	Address  address `json:"address"`
	IsActive bool
	CreateAt time.Time
	UpdateAt time.Time
}

type DeleteRequest struct {
	Email string `json:"email"`
}
