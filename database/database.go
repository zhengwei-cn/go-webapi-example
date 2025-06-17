package database

import (
	"go-webapi-example/config"
	"go-webapi-example/models"
	"go-webapi-example/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
}

// InitializeSuperAdmin 初始化超级管理员
func InitializeSuperAdmin(db *gorm.DB) error {
	// 创建用户服务实例
	userService := services.NewUserService(db)

	return userService.CreateSuperAdmin()
}
