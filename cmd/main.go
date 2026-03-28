package main

import (
	"golang-clean-architechture/internal/config"
	"golang-clean-architechture/internal/database"
	"golang-clean-architechture/internal/delivery/http/handlers"
	"golang-clean-architechture/internal/delivery/http/routes"
	"golang-clean-architechture/internal/infrastructure/persistence"
	"golang-clean-architechture/internal/infrastructure/security"
	"golang-clean-architechture/internal/usecase"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	db := database.NewMySQL(cfg)

	userRepo := persistence.NewUserGormRepository(db)
	refreshRepo := persistence.NewRefreshTokenGormRepository(db)
	passwordSvc := security.NewBcryptPasswordService()
	tokenSvc := security.NewJWTTokenService(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.JWTIssuer,
		cfg.JWTAccessExpireMinutes,
		cfg.JWTRefreshExpireDays,
	)

	authUsecase := usecase.NewAuthUsecase(db, userRepo, refreshRepo, passwordSvc, tokenSvc)
	authHandler := handlers.NewAuthHandler(authUsecase)

	departmentRepo := persistence.NewDepartmentGormRepository(db)
	departmentUsecase := usecase.NewDepartmentUseCase(db, departmentRepo)
	departmentHandler := handlers.NewDepartmentHandler(departmentUsecase)

	categoryRepo := persistence.NewCategoryGormRepository(db)
	categoryUsecase := usecase.NewCategoryUseCase(db, categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)

	app := fiber.New()
	routes.Setup(app,
		authHandler,
		departmentHandler,
		categoryHandler,
		tokenSvc,
	)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
