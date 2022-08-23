package main

import (
	"GuideGo/config"
	"GuideGo/controllers/user_controller"
	"GuideGo/repositories/user_repository"
	"GuideGo/routes"
	"GuideGo/services/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
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
