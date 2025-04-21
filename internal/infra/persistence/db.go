// internal/infra/persistence/db.go

package persistence

import (
	"go-grpc-clean/internal/adapter/persistence"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=clean port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto migrate model (เฉพาะตอน dev)
	err = db.AutoMigrate(&persistence.UserModel{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
