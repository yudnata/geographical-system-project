package points

import (
	"backend/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, h *Handler, authMiddleware fiber.Handler) {
	pts := router.Group("/points")
	pts.Post("/", authMiddleware, h.Create)
	pts.Put("/:id", authMiddleware, h.Update)
	pts.Delete("/:id", authMiddleware, h.Delete)
	pts.Get("/me", authMiddleware, h.GetMy)
	pts.Get("/", h.GetAll)

	pts.Get("/:id/blog", h.GetBlog)
	pts.Put("/:id/blog", authMiddleware, h.UpsertBlog)

	router.Post("/upload", authMiddleware, h.Upload)

	cat := router.Group("/categories")

	cat.Get("/", h.GetCategories)
	cat.Post("/", authMiddleware, middleware.RequireRole("admin"), h.CreateCategory)
	cat.Put("/:id", authMiddleware, middleware.RequireRole("admin"), h.UpdateCategory)
	cat.Delete("/:id", authMiddleware, middleware.RequireRole("admin"), h.DeleteCategory)
}
