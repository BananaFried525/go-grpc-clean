// internal/infra/persistence/db.go

package persistence

import (
	"fmt"
	"go-grpc-clean/internal/adapter/persistence"
	"go-grpc-clean/internal/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	config := config.GetDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&persistence.UserModel{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("connected to database")
	return db
}
