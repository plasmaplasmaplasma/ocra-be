package routes

import (
	"ocra/api/handlers"
	"ocra/pkg/user"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(router fiber.Router, service user.Service) {
	router.Post("/login", handlers.Login(service))
}
