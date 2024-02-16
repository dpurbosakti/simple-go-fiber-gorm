package route

import (
	"github.com/dpurbosakti/fiber-gorm/handler"
	"github.com/dpurbosakti/fiber-gorm/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Use(middlewares.Logger)
	app.Get("/users", handler.GetAllUser).Name("get_all_users")
	app.Get("/users/:id", handler.GetUserByID).Name("get_user_by_id")
	app.Post("/users", handler.CreateUser).Name("create_user")
}
