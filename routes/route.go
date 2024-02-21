package routes

import (
	"github.com/dpurbosakti/fiber-gorm/app/handlers"
	"github.com/dpurbosakti/fiber-gorm/app/queries"
	"github.com/dpurbosakti/fiber-gorm/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(app *fiber.App, db *gorm.DB) {
	queries := queries.NewQueries(db)
	handler := handlers.New(queries)

	app.Use(middlewares.Logger)

	userGroup := app.Group("/users")
	userGroup.Get("", handler.GetAllUser).Name("get_all_users")
	userGroup.Get("/:id", handler.GetUserByID).Name("get_user_by_id")
	userGroup.Post("", handler.CreateUser).Name("create_user")
	userGroup.Put("/:id", handler.UpdateUser).Name("update_user")
	userGroup.Delete("/:id", handler.DeleteUser).Name("delete_user")
}
