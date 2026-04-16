package auth

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router, h *Handler, authMiddleware fiber.Handler) {
	auth := router.Group("/auth")
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/sso", h.SSOLogin)
	auth.Put("/profile", authMiddleware, h.UpdateProfile)
}
