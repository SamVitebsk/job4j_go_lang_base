package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "ID is required")
	}

	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if strings.TrimSpace(req.Name) == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	err := s.Repository.UpdateItem(c.Context(), tracker.Item{
		Name: req.Name,
		ID:   id,
	})
	if err != nil {
		log.Errorw("s.Repository.UPdate", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
