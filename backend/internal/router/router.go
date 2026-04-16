package router

import (
	"backend/internal/modules/auth"
	"backend/internal/modules/points"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App, authH *auth.Handler, pointsH *points.Handler, authMiddleware fiber.Handler) {
	api := app.Group("/api")

	auth.RegisterRoutes(api, authH, authMiddleware)
	points.RegisterRoutes(api, pointsH)
}
