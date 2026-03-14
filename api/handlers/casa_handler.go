package handlers

import (
	"strconv"

	"ocra/api/presenter"
	"ocra/pkg/casa"

	"github.com/gofiber/fiber/v3"
)

func ListCase(service casa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		filter := casa.Filter{
			SortBy:  c.Query("sort_by"),
			SortDir: c.Query("sort_dir"),
		}
		if v := c.Query("page"); v != "" {
			if p, err := strconv.Atoi(v); err == nil {
				filter.Page = p
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

		casas, total, err := service.List(filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.CasaErrorResponse(err))
		}
		page := filter.Page
		if page < 1 {
			page = 1
		}
		return c.JSON(presenter.CasaListSuccessResponse(casas, total, page))
	}
}

func CreateCasa(service casa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req casa.CreateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(err))
		}
		result, err := service.Create(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.CasaSuccessResponse(result))
	}
}

func UpdateCasa(service casa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		var req casa.UpdateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(err))
		}
		result, err := service.Update(int64(n), req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(err))
		}
		return c.JSON(presenter.CasaSuccessResponse(result))
	}
}

func DeleteCasa(service casa.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.CasaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		if err := service.Delete(int64(n)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.CasaErrorResponse(err))
		}
		return c.JSON(&fiber.Map{"status": true, "data": nil, "error": nil})
	}
}
