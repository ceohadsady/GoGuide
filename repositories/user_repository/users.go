package user_repository

import (
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(request *User) (*User, error)
	GetUserByEmail(request string) (*User, error)
	GetAllUser() (*User, error)
	UpdateUserByEmail(email string, request *User) (*User, error) //update by email
	DeleteUserById(email string) error
	//GetUserByAge(email string,age string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateUser(request *User) (*User, error) {
	err := r.db.Create(&request).Error
	if err != nil {
		return nil, err
	}
	return request, err
}

func (r userRepository) GetUserByEmail(request string) (*User, error) {
	user := User{}
	err := r.db.Where("email", request).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
func (r userRepository) GetAllUser() (*User, error) {
	user := User{}
	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
func (r userRepository) UpdateUserByEmail(email string, request *User) (*User, error) {
	err := r.db.Model(&User{}).Where("email", email).Updates(&request).Find(&request).Error
	if err != nil {
		return nil, err
	}
	return request, err
}
func (r userRepository) DeleteUserById(email string) error {
	user := User{}
	err := r.db.Unscoped().Where("email", email).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	//auto Migrate that use to Create Table (User) or what you want
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil
	}
	return userRepository{db: db}
}
