package points

import (
	"backend/pkg/response"
	"github.com/gofiber/fiber/v3"
	"strconv"
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

	return c.Status(fiber.StatusCreated).JSON(response.Success("Titik berhasil disimpan", point))
}

func (h *Handler) GetMy(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	points, err := h.service.GetMyPoints(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal mengambil titik milik pengguna"))
	}
	return c.JSON(response.Success("Berhasil mengambil titik milik pengguna", points))
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

func (h *Handler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID tidak valid"))
	}

	var input CreatePointReq
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Format input tidak valid: " + err.Error()))
	}

	userID := c.Locals("userID").(string)
	point, err := h.service.UpdatePoint(c.Context(), userID, id, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal memperbarui data: " + err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success("Titik bangunan berhasil diperbarui", point))
}

func (h *Handler) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID tidak valid"))
	}

	userID := c.Locals("userID").(string)
	if err := h.service.DeletePoint(c.Context(), userID, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal menghapus data: " + err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success("Titik bangunan berhasil dihapus", nil))
}
