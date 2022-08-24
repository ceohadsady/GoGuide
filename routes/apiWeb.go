package routes

import (
	"GoGuide/controllers/user_controller"
	middleware "GoGuide/middleware"
	"github.com/gofiber/fiber/v2"
)

type userRoute struct {
	userCtr user_controller.UserController
}

func (r userRoute) Install(app *fiber.App) {
	adminRouter := app.Group("api/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})
	adminRouter.Post("/user/create", r.userCtr.CreateUserCrl)
	adminRouter.Post("/user/login", r.userCtr.UserLoginCrl)
	//adminRouter.Post("user/email", r.userCtr.GetUserByEmailCrl)
	adminRouter.Post("user/update", r.userCtr.UpdateUserCrl)
	adminRouter.Post("user/delete", r.userCtr.DeleteUserCrl)
	adminAuthRouter := adminRouter.Group("/auth/", middleware.NewAuthentication, func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	adminAuthRouter.Post("user/email", r.userCtr.GetUserByEmailCrl)
	//adminAuthRouter.Post("user/update", r.userCtr.UpdateUserCrl)
	//adminAuthRouter.Post("user/delete", r.userCtr.DeleteUserCrl)

}

func NewRouterAPI(userCtr user_controller.UserController) Routes {
	return userRoute{userCtr: userCtr}
}
