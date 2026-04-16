package auth

import (
	"backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Validasi gagal"))
	}
	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(response.Error(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success("Berhasil mendaftar", user))
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Validasi gagal"))
	}
	token, user, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success("Berhasil login!", fiber.Map{"token": token, "user": user}))
}

func (h *Handler) SSOLogin(c fiber.Ctx) error {
	var req SSOLoginReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Validasi SSO gagal"))
	}
	token, user, err := h.service.SSOLogin(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error("Gagal otentikasi pihak ketiga"))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success("Berhasil login otomatis via SSO", fiber.Map{"token": token, "user": user}))
}

func (h *Handler) UpdateProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var req UpdateProfileReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Validasi formulir gagal"))
	}

	user, err := h.service.UpdateProfile(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success("Profil berhasil dilengkapi!", user))
}
