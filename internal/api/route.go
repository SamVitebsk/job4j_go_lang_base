package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/items", s.CreateItem)
	route.Put("/items/:id", s.UpdateItem)
	route.Delete("/items/:id", s.DeleteItem)
	route.Get("/items", s.GetItems)
}
