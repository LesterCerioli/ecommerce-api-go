package cmd

import (
	"ecommerce-api-go/controllers"
	"ecommerce-api-go/initializers"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(services *initializers.Services) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
		}
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header format. Use 'Bearer <token>'"})
		}
		if _, err := services.AuthService.ValidateToken(c.Context(), tokenString); err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token", "reason": err.Error()})
		}
		return c.Next()
	}
}
func initializeRoutes(app *fiber.App, services *initializers.Services) {
	authController := controllers.NewAuthController(services.AuthService)

	api := app.Group("/")
	// Public auth routes
	auth := api.Group("/auth")
	auth.Post("/token", authController.GenerateToken)
    auth.Post("/token/validate", authController.ValidateToken)

}
func configureMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	db := initializers.InitialDB()
	initializers.RunMigrations(db)

	
	services := initializers.InitServices(db)

	// Start background workers
	

	app := fiber.New(fiber.Config{AppName: "Medical SAS API v1"})

	configureMiddleware(app)

	swaggerUser := os.Getenv("SWAGGER_USER")
	swaggerPass := os.Getenv("SWAGGER_PASSWORD")
	if swaggerUser == "" || swaggerPass == "" {
		log.Fatal("SWAGGER_USER and SWAGGER_PASSWORD must be set in .env file")
	}

	swaggerAuth := basicauth.New(basicauth.Config{
		Users: map[string]string{swaggerUser: swaggerPass},
		Realm: "Swagger Restricted",
	})

	app.Get("/swagger/*", swaggerAuth, swagger.HandlerDefault)
	app.Get("/docs", swaggerAuth, func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	initializeRoutes(app, services)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000
	}

	log.Printf("Server is running on http://0.0.0.0:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
