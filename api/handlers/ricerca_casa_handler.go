package handlers

import (
	"strconv"

	"ocra/api/presenter"
	"ocra/pkg/ricercacasa"

	"github.com/gofiber/fiber/v3"
)

func ListRicercaCase(service ricercacasa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		filter := ricercacasa.Filter{}

		if v := c.Query("page"); v != "" {
			if p, err := strconv.Atoi(v); err == nil {
				filter.Page = p
			}
		}
		if v := c.Query("tempo_di_acquisto_from"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 32); err == nil {
				i := int32(n)
				filter.TempoDiAcquistoFrom = &i
			}
		}
		if v := c.Query("tempo_di_acquisto_to"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 32); err == nil {
				i := int32(n)
				filter.TempoDiAcquistoTo = &i
			}
		}
		if v := c.Query("budget_from"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.BudgetFrom = &f
			}
		}
		if v := c.Query("budget_to"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.BudgetTo = &f
			}
		}
		if v := c.Query("percentuale_mutuo_from"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.PercentualeMutuoFrom = &f
			}
		}
		if v := c.Query("percentuale_mutuo_to"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.PercentualeMutuoTo = &f
			}
		}
		if v := c.Query("liquidita_from"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.LiquiditaFrom = &f
			}
		}
		if v := c.Query("liquidita_to"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				filter.LiquiditaTo = &f
			}
		}
		if v := c.Query("cliente_id"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				id := int64(n)
				filter.ClienteID = &id
			}
		}
		if v := c.Query("piano"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 16); err == nil {
				i := int16(n)
				filter.Piano = &i
			}
		}
		if v := c.Query("numero_di_locali"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 16); err == nil {
				i := int16(n)
				filter.NumeroDiLocali = &i
			}
		}
		if v := c.Query("numero_di_camere"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 16); err == nil {
				i := int16(n)
				filter.NumeroDiCamere = &i
			}
		}
		if v := c.Query("numero_di_bagni"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 16); err == nil {
				i := int16(n)
				filter.NumeroDiBagni = &i
			}
		}
		if v := c.Query("balcone"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.Balcone = &b
			}
		}
		if v := c.Query("terrazzo"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.Terrazzo = &b
			}
		}
		if v := c.Query("giardino"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.Giardino = &b
			}
		}
		if v := c.Query("zona_id"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				id := int64(n)
				filter.ZonaID = &id
			}
		}

		results, total, err := service.List(filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		page := filter.Page
		if page < 1 {
			page = 1
		}
		return c.JSON(presenter.RicercaCasaListSuccessResponse(results, total, page))
	}
}

func CreateRicercaCasa(service ricercacasa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req ricercacasa.CreateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		result, err := service.Create(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.RicercaCasaSuccessResponse(result))
	}
}

func UpdateRicercaCasa(service ricercacasa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		var req ricercacasa.UpdateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		result, err := service.Update(int64(n), req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		return c.JSON(presenter.RicercaCasaSuccessResponse(result))
	}
}

func DeleteRicercaCasa(service ricercacasa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.RicercaCasaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		if err := service.Delete(int64(n)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.RicercaCasaErrorResponse(err))
		}
		return c.JSON(&fiber.Map{"status": true, "data": nil, "error": nil})
	}
}
