package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type ItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetItemsResponse struct {
	Items []ItemRequest `json:"items"`
}

func (s *Server) GetItems(c *fiber.Ctx) error {
	var items []tracker.Item
	var err error
	name := c.Query("name")
	if name == "" {
		items, err = s.Repository.List(c.Context())
	} else {
		items, err = s.Repository.FindByNameLike(c.Context(), name)
	}

	if err != nil {
		log.Errorw("s.Repository.List", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{Items: res})
}
