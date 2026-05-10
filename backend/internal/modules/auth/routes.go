package auth

import (
	"backend/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, h *Handler, authMiddleware fiber.Handler) {
	auth := router.Group("/auth")

	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/sso", h.SSOLogin)
	auth.Get("/config", h.GetConfig)

	auth.Get("/me", authMiddleware, h.GetMe)
	auth.Put("/profile", authMiddleware, h.UpdateProfile)
	auth.Put("/password", authMiddleware, h.UpdatePassword)

	auth.Get("/users", authMiddleware, middleware.RequireRole("admin"), h.GetUsers)
}
