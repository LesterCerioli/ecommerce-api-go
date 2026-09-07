package initializers

import (
	"context"
	"ecommerce-api-go/shared/services"
	"log"

	"gorm.io/gorm"
)

type Services struct {
	AuthService *services.AuthService
}

func InitServices(db *gorm.DB) *Services {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	log.Println("Initializing services...")

	authService, err := services.NewAuthService(sqlDB)
	if err != nil {
		log.Fatalf("Failed to initialize AuthService: %v", err)
	}

	if err := authService.EnsureTable(context.Background()); err != nil {
		log.Fatalf("Failed to ensure auth_tokens table: %v", err)
	}

	return &Services{
		AuthService: authService,
	}
}
