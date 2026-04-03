package database

import (
	"golang-clean-architechture/internal/config"
	"golang-clean-architechture/internal/infrastructure/persistence/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	if err := db.AutoMigrate(
		&models.UserModel{},
		&models.RefreshTokenModel{},
		&models.DepartmentModel{},
		&models.CategoryModel{},
		&models.SubjectModel{},
		&models.SubjectItemModel{},
		&models.SubjectTypeModel{},
	); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	return db
}
