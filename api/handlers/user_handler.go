package handlers

import (
	"fmt"
	"ocra/api/presenter"
	"ocra/pkg/entities"
	"ocra/pkg/user"

	"github.com/gofiber/fiber/v3"

	"strings"
)

func Login(service user.Service) fiber.Handler {
	fmt.Println("received login req")
	return func(c fiber.Ctx) error {
		var req entities.LoginRequest

		if err := c.Bind().Body(&req); err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
			return c.JSON(presenter.UserErrorResponse(fiber.NewError(fiber.StatusBadRequest, "username and password are required")))
		}

		loginUser, err := service.Login(req.Email, req.Password)
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(loginUser))
	}
}
