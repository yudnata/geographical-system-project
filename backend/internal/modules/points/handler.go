package points

import (
	"backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var input CreatePointReq

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Invalid input format: " + err.Error()))
	}

	if input.Name == "" || input.Latitude == 0 || input.Longitude == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Name, Latitude, dan Longitude wajib diisi"))
	}

	ownerID := c.Locals("userID")
	if ownerID == nil {
		ownerID = "00000000-0000-0000-0000-000000000000"
	}

	point, err := h.service.CreatePoint(c.Context(), ownerID.(string), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal menyimpan data: " + err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success("Titik bangunan berhasil ditambahkan", point))
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	points, err := h.service.GetAllPoints(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal memuat titik peta: " + err.Error()))
	}

	if points == nil {
		points = []GeoPoint{}
	}

	return c.Status(fiber.StatusOK).JSON(response.Success("Berhasil mendapatkan data peta", points))
}
