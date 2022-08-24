package user_controller

import (
	"GoGuide/controllers"
	"GoGuide/logs"
	"GoGuide/services/user_service"
	"GoGuide/validation"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUserCrl(ctx *fiber.Ctx) error
	UserLoginCrl(ctx *fiber.Ctx) error
	GetUserByEmailCrl(ctx *fiber.Ctx) error
	UpdateUserCrl(ctx *fiber.Ctx) error
	DeleteUserCrl(ctx *fiber.Ctx) error
}

type userController struct {
	userSrv user_service.UserService
}

func (c userController) CreateUserCrl(ctx *fiber.Ctx) error {
	//ຮັບຄ່າ UserCreateRequest ມາ validate
	validate := new(user_service.UserCreateRequest)
	err := ctx.BodyParser(validate)              // harm h pen pointer
	errs := validation.ValidateStruct(*validate) // ya leum h pen pointer
	if errs != nil {
		logs.Error(errs)
		return controllers.NewErrorValidate(ctx, errs[0].Error) // errs[0].Error return hai flutter jut karn theua la 1 error
	}
	request := user_service.UserCreateRequest{}
	err = ctx.BodyParser(&request)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	userCreate, err := c.userSrv.CreateUserSrv(&request)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewCreateSuccessResponse(ctx, &userCreate)
}

func (c userController) UserLoginCrl(ctx *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctx.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	userLogin, err := c.userSrv.UserLoginSrv(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, &userLogin)
}

func (c userController) GetUserByEmailCrl(ctx *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctx.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	userByMail, err := c.userSrv.GetUserByEmailSrv(user.Email)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, &userByMail)
}
func (c userController) UpdateUserCrl(ctx *fiber.Ctx) error {
	user := user_service.UserUpdateRequest{}
	err := ctx.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
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
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, &userRepo)
}
func (c userController) DeleteUserCrl(ctx *fiber.Ctx) error {
	user := user_service.UserLoginRequest{}
	err := ctx.BodyParser(&user)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	err = c.userSrv.DeleteUserSrv(user.Email)
	if err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, "DELETED")
}

func NewControllerUser(userSrv user_service.UserService) UserController {
	return userController{userSrv: userSrv}
}
