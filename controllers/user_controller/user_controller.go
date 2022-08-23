package user_controller

import (
	"GuideGo/controllers"
	"GuideGo/logs"
	"GuideGo/services/user_service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUserCrl(ctc *fiber.Ctx) error
	UserLoginCrl(ctc *fiber.Ctx) error
	GetUserByEmailCrl(ctc *fiber.Ctx) error
	UpdateUserCrl(ctc *fiber.Ctx) error
	DeleteUserCrl(ctc *fiber.Ctx) error
}

type userController struct {
	userSrv user_service.UserService
}

func (c userController) CreateUserCrl(ctc *fiber.Ctx) error {
	user := user_service.UserCreateRequest{}
	err := ctc.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	userCreate, err := c.userSrv.CreateUserSrv(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	return controllers.NewCreateSuccessResponse(ctc, &userCreate)
}
func (c userController) UserLoginCrl(ctc *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctc.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	userLogin, err := c.userSrv.UserLoginSrv(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	return controllers.NewSuccessResponse(ctc, &userLogin)
}
func (c userController) GetUserByEmailCrl(ctc *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctc.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	userByMail, err := c.userSrv.GetUserByEmailSrv(user.Email)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	return controllers.NewSuccessResponse(ctc, &userByMail)
}
func (c userController) UpdateUserCrl(ctc *fiber.Ctx) error {
	user := user_service.UserUpdateRequest{}
	err := ctc.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	userRepo, err := c.userSrv.UpdateUserSrv(user.Email, user_service.UserUpdateRequest{
		Name:     user.Name,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	})
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	return controllers.NewSuccessResponse(ctc, &userRepo)
}
func (c userController) DeleteUserCrl(ctc *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctc.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	err = c.userSrv.DeleteUserSrv(user.Email)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctc, err)
	}
	return controllers.NewSuccessResponse(ctc, "DELETED")
}

func NewControllerUser(userSrv user_service.UserService) UserController {
	return userController{userSrv: userSrv}
}
