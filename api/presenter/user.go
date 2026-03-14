package presenter

import (
	"ocra/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	ID       int64
	Email    string
	Username string
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
