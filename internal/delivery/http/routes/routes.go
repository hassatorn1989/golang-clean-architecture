package routes

import (
	"golang-clean-architechture/internal/delivery/http/handlers"
	"golang-clean-architechture/internal/domain/services"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App,
	authHandler *handlers.AuthHandler,
	departmentHandler *handlers.DepartmentHandler,
	categoryHandler *handlers.CategoryHandler,
	tokenSvc services.TokenService) {
	api := app.Group("/api")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "ok"})
	})

	// auth := api.Group("/auth")
	// auth.Post("/register", authHandler.Register)
	// auth.Post("/login", authHandler.Login)
	// auth.Post("/refresh", authHandler.Refresh)
	// auth.Post("/logout", authHandler.Logout)

	// protected := api.Group("/user", middleware.AuthMiddleware(tokenSvc))
	// protected.Get("/me", authHandler.Me)
	// , middleware.AuthMiddleware(tokenSvc)

	department := api.Group("/departments")
	department.Get("/", departmentHandler.GetAll)
	department.Get("/pagination", departmentHandler.GetWithPagination)
	department.Get("/:id", departmentHandler.GetByID)
	department.Post("/", departmentHandler.Create)
	department.Put("/:id", departmentHandler.Update)
	department.Delete("/:id", departmentHandler.Delete)

	category := api.Group("/categories")
	category.Get("/", categoryHandler.GetAll)
	category.Get("/pagination", categoryHandler.GetWithPagination)
	category.Get("/:id", categoryHandler.GetByID)
	category.Post("/", categoryHandler.Create)
	category.Put("/:id", categoryHandler.Update)
	category.Delete("/:id", categoryHandler.Delete)
}
