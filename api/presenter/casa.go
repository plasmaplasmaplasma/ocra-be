package presenter

import (
	"ocra/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type CasaResponse struct {
	ID             int64        `json:"id"`
	Piano          *int16       `json:"piano"`
	NumeroDiLocali *int16       `json:"numero_di_locali"`
	NumeroDiCamere *int16       `json:"numero_di_camere"`
	NumeroDiBagni  *int16       `json:"numero_di_bagni"`
	Balcone        *bool        `json:"balcone"`
	Terrazzo       *bool        `json:"terrazzo"`
	Giardino       *bool        `json:"giardino"`
	ZonaID         int64        `json:"zona_id"`
	Zona           ZonaResponse `json:"zona"`
}

func ToCasaResponse(c *entities.House) CasaResponse {
	return CasaResponse{
		ID:             c.ID,
		Piano:          c.Piano,
		NumeroDiLocali: c.NumeroDiLocali,
		NumeroDiCamere: c.NumeroDiCamere,
		NumeroDiBagni:  c.NumeroDiBagni,
		Balcone:        c.Balcone,
		Terrazzo:       c.Terrazzo,
		Giardino:       c.Giardino,
		ZonaID:         c.ZonaID,
		Zona:           ToZonaResponse(&c.Zona),
	}
}

func CasaSuccessResponse(data *entities.House) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   ToCasaResponse(data),
		"error":  nil,
	}
}

func CasaListSuccessResponse(data []entities.House, total int64, page int) *fiber.Map {
	responses := make([]CasaResponse, len(data))
	for i, c := range data {
		responses[i] = ToCasaResponse(&c)
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

func CasaErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
