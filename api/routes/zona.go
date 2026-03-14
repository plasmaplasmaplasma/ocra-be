package routes

import (
	"ocra/api/handlers"
	"ocra/pkg/zona"

	"github.com/gofiber/fiber/v3"
)

func ZonaRouter(router fiber.Router, service zona.Service) {
	router.Get("/read", handlers.ListZone(service))
	router.Post("/create", handlers.CreateZona(service))
	router.Put("/update/:id", handlers.UpdateZona(service))
	router.Delete("/delete/:id", handlers.DeleteZona(service))
}
