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
	subjectHandler *handlers.SubjectHandler,
	subjectTypeHandler *handlers.SubjectTypeHandler,
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

	subject := api.Group("/subjects")
	subject.Get("/", subjectHandler.GetAll)
	subject.Get("/pagination", subjectHandler.GetWithPagination)
	subject.Get("/:id", subjectHandler.GetByID)
	subject.Post("/", subjectHandler.Create)
	subject.Put("/:id", subjectHandler.Update)
	subject.Delete("/:id", subjectHandler.Delete)

	subjectType := api.Group("/subject-types")
	subjectType.Get("/", subjectTypeHandler.GetAll)
	subjectType.Get("/pagination", subjectTypeHandler.GetWithPagination)
	subjectType.Get("/:id", subjectTypeHandler.GetByID)
	subjectType.Post("/", subjectTypeHandler.Create)
	subjectType.Put("/:id", subjectTypeHandler.Update)
	subjectType.Delete("/:id", subjectTypeHandler.Delete)
}
