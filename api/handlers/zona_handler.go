package handlers

import (
	"strconv"

	"ocra/api/presenter"
	"ocra/pkg/zona"

	"github.com/gofiber/fiber/v3"
)

func ListZone(service zona.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		page := 1
		if v := c.Query("page"); v != "" {
			if p, err := strconv.Atoi(v); err == nil && p > 0 {
				page = p
			}
		}
		zones, total, err := service.List(page)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ZonaErrorResponse(err))
		}
		return c.JSON(presenter.ZonaListSuccessResponse(zones, total, page))
	}
}

func CreateZona(service zona.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req zona.CreateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(err))
		}
		result, err := service.Create(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.ZonaSuccessResponse(result))
	}
}

func UpdateZona(service zona.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		var req zona.UpdateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(err))
		}
		result, err := service.Update(int64(n), req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(err))
		}
		return c.JSON(presenter.ZonaSuccessResponse(result))
	}
}

func DeleteZona(service zona.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ZonaErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		if err := service.Delete(int64(n)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ZonaErrorResponse(err))
		}
		return c.JSON(&fiber.Map{"status": true, "data": nil, "error": nil})
	}
}
