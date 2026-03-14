package routes

import (
	"ocra/api/handlers"
	"ocra/pkg/cliente"

	"github.com/gofiber/fiber/v3"
)

func ClienteRouter(router fiber.Router, service cliente.Service) {
	router.Get("/read", handlers.ListClienti(service))
	router.Post("/create", handlers.CreateCliente(service))
	router.Put("/update/:id", handlers.UpdateCliente(service))
	router.Delete("/delete/:id", handlers.DeleteCliente(service))
}
