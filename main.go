package main

import (
	"GoGuide/config"
	"GoGuide/controllers/user_controller"
	"GoGuide/repositories/user_repository"
	"GoGuide/routes"
	"GoGuide/services/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Vientiane",
	}))
	app.Use(cors.New())
	connectDB, err := config.ConnectDB()
	if err != nil {
		return
	}
	userRepository := user_repository.NewUserRepository(connectDB)
	userService := user_service.NewUserService(userRepository)
	userController := user_controller.NewControllerUser(userService)
	routerApi := routes.NewRouterAPI(userController)
	routerApi.Install(app)
	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
