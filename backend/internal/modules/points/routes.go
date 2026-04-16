package points

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router, h *Handler) {
	pts := router.Group("/points")
	pts.Post("/", h.Create)
	pts.Get("/", h.GetAll)
}
