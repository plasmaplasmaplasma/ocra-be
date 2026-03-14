package presenter

import (
	"ocra/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type RicercaCasaResponse struct {
	ID               int64           `json:"id"`
	TempoDiAcquisto  int32           `json:"tempo_di_acquisto"`
	Budget           float64         `json:"budget"`
	PercentualeMutuo float64         `json:"percentuale_mutuo"`
	Liquidita        float64         `json:"liquidita"`
	ClienteID        int64           `json:"cliente_id"`
	Cliente          ClienteResponse `json:"cliente"`
	CasaID           int64           `json:"casa_id"`
	Casa             CasaResponse    `json:"casa"`
}

func ToRicercaCasaResponse(rc *entities.SearchHouse) RicercaCasaResponse {
	return RicercaCasaResponse{
		ID:               rc.ID,
		TempoDiAcquisto:  rc.TempoDiAcquisto,
		Budget:           rc.Budget,
		PercentualeMutuo: rc.PercentualeMutuo,
		Liquidita:        rc.Liquidita,
		ClienteID:        rc.ClienteID,
		Cliente:          ToClienteResponse(&rc.Cliente),
		CasaID:           rc.CasaID,
		Casa:             ToCasaResponse(&rc.Casa),
	}
}

func RicercaCasaSuccessResponse(data *entities.SearchHouse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   ToRicercaCasaResponse(data),
		"error":  nil,
	}
}

func RicercaCasaListSuccessResponse(data []entities.SearchHouse, total int64, page int) *fiber.Map {
	responses := make([]RicercaCasaResponse, len(data))
	for i, rc := range data {
		responses[i] = ToRicercaCasaResponse(&rc)
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

func RicercaCasaErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
