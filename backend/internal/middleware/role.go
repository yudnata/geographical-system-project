package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func RequireRole(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := c.Locals("role")
		if userRole == nil {
			return c.Status(403).JSON(fiber.Map{"success": false, "message": "Akses ditolak: Peran tidak ditemukan"})
		}

		roleStr, ok := userRole.(string)
		if !ok {
			return c.Status(403).JSON(fiber.Map{"success": false, "message": "Akses ditolak: Format peran tidak valid"})
		}

		for _, role := range roles {
			if roleStr == role {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{"success": false, "message": "Akses ditolak: Anda tidak memiliki izin untuk tindakan ini"})
	}
}
