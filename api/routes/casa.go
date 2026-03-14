package routes

import (
	"ocra/api/handlers"
	"ocra/pkg/casa"

	"github.com/gofiber/fiber/v3"
)

func CasaRouter(router fiber.Router, service casa.Service) {
	router.Get("/read", handlers.ListCase(service))
	router.Post("/create", handlers.CreateCasa(service))
	router.Put("/update/:id", handlers.UpdateCasa(service))
	router.Delete("/delete/:id", handlers.DeleteCasa(service))
}
