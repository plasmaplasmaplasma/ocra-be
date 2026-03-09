package main

import (
	"fmt"
	"ocra/api/routes"
	"ocra/database"
	"ocra/pkg/user"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found")
		return
	}

	dbSchema := os.Getenv("DB_SCHEMA")
	if dbSchema == "" {
		fmt.Println("Warning: DB_SCHEMA not set in .env file")
		return
	}

	database := database.Setup()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	api := app.Group("/api")

	userApi := api.Group("/users")

	userTable := database.Table(dbSchema + ".users")
	userRepo := user.NewRepository(userTable)
	userService := user.NewService(userRepo)
	routes.UserRouter(userApi, userService)

	fmt.Printf("Server starting on port %s\n", appPort)
	if err := app.Listen(":" + appPort); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
