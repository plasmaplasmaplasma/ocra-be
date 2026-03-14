package main

import (
	"fmt"
	"ocra/api/routes"
	"ocra/database"
	"ocra/pkg/casa"
	"ocra/pkg/cliente"
	"ocra/pkg/ricercacasa"
	user "ocra/pkg/utente"
	"ocra/pkg/zona"
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

	db := database.Setup()

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
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	routes.UserRouter(userApi, userService)

	clienteApi := api.Group("/clients")
	clienteRepo := cliente.NewRepository(db)
	clienteService := cliente.NewService(clienteRepo)
	routes.ClienteRouter(clienteApi, clienteService)

	zonaApi := api.Group("/zones")
	zonaRepo := zona.NewRepository(db)
	zonaService := zona.NewService(zonaRepo)
	routes.ZonaRouter(zonaApi, zonaService)

	casaApi := api.Group("/houses")
	casaRepo := casa.NewRepository(db)
	casaService := casa.NewService(casaRepo)
	routes.CasaRouter(casaApi, casaService)

	ricercaCasaApi := api.Group("/search_houses")
	ricercaCasaRepo := ricercacasa.NewRepository(db)
	ricercaCasaService := ricercacasa.NewService(ricercaCasaRepo)
	routes.RicercaCasaRouter(ricercaCasaApi, ricercaCasaService)

	fmt.Printf("Server starting on port %s\n", appPort)
	if err := app.Listen(":" + appPort); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
