package points

import (
	"backend/pkg/response"
	"backend/pkg/upload"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type Handler struct {
	service  *Service
	uploader *upload.CloudinaryService
}

func NewHandler(service *Service, uploader *upload.CloudinaryService) *Handler {
	return &Handler{service: service, uploader: uploader}
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
		points = []MapPoint{}
	}

	return c.Status(fiber.StatusOK).JSON(response.Success("Berhasil mendapatkan data peta", points))
}

func (h *Handler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID tidak valid"))
	}

	var input UpdatePointReq
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

func (h *Handler) GetCategories(c fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal mengambil kategori"))
	}
	if categories == nil {
		categories = []Category{}
	}
	return c.Status(fiber.StatusOK).JSON(response.Success("Berhasil mengambil kategori", categories))
}

func (h *Handler) CreateCategory(c fiber.Ctx) error {
	var input struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	if err := c.Bind().Body(&input); err != nil || input.Name == "" || input.Icon == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Nama dan icon wajib diisi"))
	}
	cat, err := h.service.CreateCategory(input.Name, input.Icon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal membuat kategori"))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success("Kategori berhasil dibuat", cat))
}

func (h *Handler) UpdateCategory(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID tidak valid"))
	}
	var input struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	if err := c.Bind().Body(&input); err != nil || input.Name == "" || input.Icon == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Nama dan icon wajib diisi"))
	}
	cat, err := h.service.UpdateCategory(id, input.Name, input.Icon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal memperbarui kategori"))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success("Kategori berhasil diperbarui", cat))
}

func (h *Handler) DeleteCategory(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID tidak valid"))
	}
	if err := h.service.DeleteCategory(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal menghapus kategori"))
	}
	return c.Status(fiber.StatusOK).JSON(response.Success("Kategori berhasil dihapus", nil))
}

func (h *Handler) GetBlog(c fiber.Ctx) error {
	pointID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID Titik tidak valid"))
	}
	blog, err := h.service.GetBlog(c.Context(), pointID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("Blog tidak ditemukan"))
	}
	return c.JSON(response.Success("Berhasil mengambil data blog", blog))
}

func (h *Handler) UpsertBlog(c fiber.Ctx) error {
	pointID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("ID Titik tidak valid"))
	}

	var input UpsertBlogReq
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Format input tidak valid"))
	}

	userID := c.Locals("userID").(string)
	role := c.Locals("role").(string)
	point, err := h.service.repo.GetByID(c.Context(), pointID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("Titik tidak ditemukan"))
	}
	if role != "admin" && point.OwnerID != userID {
		return c.Status(fiber.StatusForbidden).JSON(response.Error("Anda tidak memiliki izin untuk mengulas titik ini"))
	}

	blog, err := h.service.UpsertBlog(c.Context(), pointID, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal menyimpan ulasan: " + err.Error()))
	}

	return c.JSON(response.Success("Ulasan berhasil disimpan", blog))
}

func (h *Handler) Upload(c fiber.Ctx) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("File tidak ditemukan"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal membuka file"))
	}
	defer file.Close()

	folder := c.FormValue("folder", "points")
	url, err := h.uploader.UploadImage(c.Context(), file, folder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error("Gagal mengunggah gambar: " + err.Error()))
	}

	return c.JSON(response.Success("Gambar berhasil diunggah", fiber.Map{"url": url}))
}
