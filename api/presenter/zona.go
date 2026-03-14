package presenter

import (
	"ocra/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type ZonaResponse struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

func ToZonaResponse(z *entities.Zone) ZonaResponse {
	return ZonaResponse{ID: z.ID, Nome: z.Nome}
}

func ZonaSuccessResponse(data *entities.Zone) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   ToZonaResponse(data),
		"error":  nil,
	}
}

func ZonaListSuccessResponse(data []entities.Zone, total int64, page int) *fiber.Map {
	responses := make([]ZonaResponse, len(data))
	for i, z := range data {
		responses[i] = ToZonaResponse(&z)
	}
	return &fiber.Map{
		"status":   true,
		"data":     responses,
		"total":    total,
		"page":     page,
		"per_page": 20,
		"error":    nil,
	}
}

func ZonaErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
