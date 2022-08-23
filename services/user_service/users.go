package user_service

import (
	"GuideGo/logs"
	"GuideGo/repositories/user_repository"
	"GuideGo/security"
	"errors"
	"time"
)

type UserService interface {
	CreateUserSrv(request *UserCreateRequest) (*UserCreateResponse, error)
	UserLoginSrv(request *UserLoginRequest) (*UserLoginResponse, error)
	GetUserByEmailSrv(request string) (*UserResponse, error)
	UpdateUserSrv(email string, request UserUpdateRequest) (*UserUpdateResponse, error)
	DeleteUserSrv(email string) error
}

type userRepository struct {
	userRepo user_repository.UsersRepository
}

func (s userRepository) CreateUserSrv(request *UserCreateRequest) (*UserCreateResponse, error) {
	//if request.FullName == "" || request.Email == "" || request.Phone == 0 || request.Name == "" || request.Password == "" {
	//	logs.Error("ENTER_ALL_FIELD")
	//	return nil, errors.New("ENTER_ALL_FIELD")
	//}
	getUser, _ := s.userRepo.GetUserByEmail(request.Email)
	if getUser.Email == request.Email {
		logs.Error("ALREADY_EXIST_USER")
		return nil, errors.New("ALREADY_EXIST_USER")
	}
	newEncryptPassword, _ := security.NewEncryptPassword(request.Password)

	userRepoCreate := user_repository.User{
		Name:     request.Name,
		FullName: request.FullName,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: newEncryptPassword,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		IsActive: true,
	}
	createUser, _ := s.userRepo.CreateUser(&userRepoCreate)
	userResponse := UserCreateResponse{
		Name:     createUser.Name,
		FullName: createUser.FullName,
		Email:    createUser.Email,
		Phone:    createUser.Phone,
		Password: createUser.Password,
		CreateAt: createUser.CreateAt,
		UpdateAt: createUser.UpdateAt,
		IsActive: createUser.IsActive,
	}
	return &userResponse, nil
}
func (s userRepository) UserLoginSrv(request *UserLoginRequest) (*UserLoginResponse, error) {
	if request.Email == "" || request.Password == "" {
		logs.Error("INVALID_EMAIL_AND_PASSWORD")
		return nil, errors.New("INVALID_EMAIL_AND_PASSWORD")
	}
	getUser, _ := s.userRepo.GetUserByEmail(request.Email)
	if getUser.Email != request.Email {
		logs.Error("INVALID_EMAIL")
		return nil, errors.New("INVALID_EMAIL")
	}
	err := security.VerifyPassword(getUser.Password, request.Password)
	if err != nil {
		logs.Error("INVALID_PASSWORD")
		return nil, errors.New("INVALID_PASSWORD")
	}
	userLoginRes := UserLoginResponse{
		Name:     getUser.Name,
		FullName: getUser.FullName,
		Email:    getUser.Email,
		Phone:    getUser.Phone,
		IsActive: getUser.IsActive,
		CreateAt: getUser.CreatedAt,
	}
	return &userLoginRes, nil
}
func (s userRepository) GetUserByEmailSrv(request string) (*UserResponse, error) {
	getUser, _ := s.userRepo.GetUserByEmail(request)
	if getUser.Email == "" {
		logs.Error("USER_NOT_FOUND")
		return nil, errors.New("USER_NOT_FOUND")
	}
	userResponse := UserResponse{
		Name:     getUser.Name,
		FullName: getUser.FullName,
		Email:    getUser.Email,
		Phone:    getUser.Phone,
		IsActive: getUser.IsActive,
		CreateAt: getUser.CreatedAt,
		UpdateAt: getUser.UpdatedAt,
	}
	return &userResponse, nil
}
func (s userRepository) UpdateUserSrv(email string, request UserUpdateRequest) (*UserUpdateResponse, error) {
	getUser, _ := s.userRepo.GetUserByEmail(email)
	if getUser.Email == "" {
		logs.Error("USER_NOT_FOUND")
		return nil, errors.New("USER_NOT_FOUND")
	}
	//if getUser.Email == request.Email {
	//	logs.Error("THIS_USERNAME_ALREADY_USED")
	//	return nil, errors.New("THIS_USERNAME_ALREADY_USED")
	//}
	userRepo := user_repository.User{
		Name:     request.Name,
		FullName: request.FullName,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
		UpdateAt: time.Now(),
	}
	userCreate, _ := s.userRepo.UpdateUserByEmail(email, &userRepo)

	userCreateRes := UserUpdateResponse{
		Name:     userCreate.Name,
		FullName: userCreate.FullName,
		Email:    userCreate.Email,
		Phone:    userCreate.Phone,
		Password: userCreate.Password,
		UpdateAt: userCreate.UpdateAt,
	}
	return &userCreateRes, nil
}
func (s userRepository) DeleteUserSrv(email string) error {
	getUser, err := s.userRepo.GetUserByEmail(email)
	if getUser.Email != email {
		logs.Error("INVALID_USER")
		return errors.New("INVALID_USER")
	}
	_ = s.userRepo.DeleteUserById(email)
	return err
}

func NewUserService(userRepo user_repository.UsersRepository) UserService {
	return &userRepository{userRepo: userRepo}

}
