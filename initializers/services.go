package initializers

import (
	"database/sql"

	"ecommerce-api-go/shared/services"
)

type Services struct {
	AuthService *services.AuthTokenService
}

func InitService(db *sql.DB) *Services {
	authService, _ := services.NewAuthTokenService(db)

	return &Services{
		AuthService: authService,
	}
}
