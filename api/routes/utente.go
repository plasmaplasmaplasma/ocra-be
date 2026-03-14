package routes

import (
	"ocra/api/handlers"
	user "ocra/pkg/utente"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(router fiber.Router, service user.Service) {
	router.Post("/login", handlers.Login(service))
}
