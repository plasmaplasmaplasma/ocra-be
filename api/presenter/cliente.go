package presenter

import (
	"ocra/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type ClienteResponse struct {
	ID                 int64   `json:"id"`
	Nome               *string `json:"nome"`
	Cognome            *string `json:"cognome"`
	NumeroDiTelefono   *string `json:"numero_di_telefono"`
	Email              *string `json:"email"`
	Acquista           bool    `json:"acquista"`
	Vende              bool    `json:"vende"`
	VendePerAcquistare bool    `json:"vende_per_acquistare"`
}

func ToClienteResponse(c *entities.Client) ClienteResponse {
	return ClienteResponse{
		ID:                 c.ID,
		Nome:               c.Nome,
		Cognome:            c.Cognome,
		NumeroDiTelefono:   c.NumeroDiTelefono,
		Email:              c.Email,
		Acquista:           c.Acquista,
		Vende:              c.Vende,
		VendePerAcquistare: c.VendePerAcquistare,
	}
}

func ClienteSuccessResponse(data *entities.Client) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   ToClienteResponse(data),
		"error":  nil,
	}
}

func ClienteListSuccessResponse(data []entities.Client, total int64, page int) *fiber.Map {
	responses := make([]ClienteResponse, len(data))
	for i, c := range data {
		responses[i] = ToClienteResponse(&c)
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

func ClienteErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
