package routes

import (
	"ocra/api/handlers"
	"ocra/pkg/ricercacasa"

	"github.com/gofiber/fiber/v3"
)

func RicercaCasaRouter(router fiber.Router, service ricercacasa.Service) {
	router.Get("/read", handlers.ListRicercaCase(service))
	router.Post("/create", handlers.CreateRicercaCasa(service))
	router.Put("/update/:id", handlers.UpdateRicercaCasa(service))
	router.Delete("/delete/:id", handlers.DeleteRicercaCasa(service))
}
