package handlers

import (
	"strconv"

	"ocra/api/presenter"
	"ocra/pkg/cliente"

	"github.com/gofiber/fiber/v3"
)

func ListClienti(service cliente.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		filter := cliente.Filter{
			SortBy:  c.Query("sort_by"),
			SortDir: c.Query("sort_dir"),
		}
		if v := c.Query("page"); v != "" {
			if p, err := strconv.Atoi(v); err == nil {
				filter.Page = p
			}
		}
		if v := c.Query("acquista"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.Acquista = &b
			}
		}
		if v := c.Query("vende"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.Vende = &b
			}
		}
		if v := c.Query("vende_per_acquistare"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				filter.VendePerAcquistare = &b
			}
		}
		if v := c.Query("zona_id"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				id := int64(n)
				filter.ZonaID = &id
			}
		}

		clientes, total, err := service.List(filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ClienteErrorResponse(err))
		}
		page := filter.Page
		if page < 1 {
			page = 1
		}
		return c.JSON(presenter.ClienteListSuccessResponse(clientes, total, page))
	}
}

func CreateCliente(service cliente.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req cliente.CreateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(err))
		}
		result, err := service.Create(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.ClienteSuccessResponse(result))
	}
}

func UpdateCliente(service cliente.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		var req cliente.UpdateRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(err))
		}
		result, err := service.Update(int64(n), req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(err))
		}
		return c.JSON(presenter.ClienteSuccessResponse(result))
	}
}

func DeleteCliente(service cliente.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		n, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ClienteErrorResponse(
				fiber.NewError(fiber.StatusBadRequest, "id non valido"),
			))
		}
		if err := service.Delete(int64(n)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ClienteErrorResponse(err))
		}
		return c.JSON(&fiber.Map{"status": true, "data": nil, "error": nil})
	}
}
