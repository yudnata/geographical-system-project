package points

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Category Services
func (s *Service) GetAllCategories() ([]Category, error) {
	return s.repo.GetAllCategories(context.Background())
}

func (s *Service) CreateCategory(name, icon string) (*Category, error) {
	return s.repo.CreateCategory(context.Background(), name, icon)
}

func (s *Service) UpdateCategory(id int, name, icon string) (*Category, error) {
	return s.repo.UpdateCategory(context.Background(), id, name, icon)
}

func (s *Service) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(context.Background(), id)
}

func (s *Service) CreatePoint(ctx context.Context, ownerID string, input CreatePointReq) (*MapPoint, error) {
	status := "draft"
	if input.Status != "" {
		status = input.Status
	}

	point := &MapPoint{
		CategoryID:   &input.CategoryID,
		Name:         input.Name,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		Address:      input.Address,
		OwnerID:      ownerID,
		TahunBerdiri: &input.TahunBerdiri,
		Description:  input.Description,
		CoverImage:   &input.CoverImage,
		Status:       status,
	}

	if err := s.repo.Create(ctx, point); err != nil {
		return nil, err
	}
	return point, nil
}

func (s *Service) GetAllPoints(ctx context.Context) ([]MapPoint, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetPublicPoints(ctx context.Context) ([]MapPoint, error) {
	return s.repo.GetPublic(ctx)
}

func (s *Service) GetMyPoints(ctx context.Context, ownerID string) ([]MapPoint, error) {
	return s.repo.GetMyPoints(ctx, ownerID)
}

func (s *Service) GetPending(ctx context.Context) ([]MapPoint, error) {
	return s.repo.GetPending(ctx)
}

func (s *Service) UpdatePoint(ctx context.Context, userID string, id int, req UpdatePointReq) (*MapPoint, error) {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("Titik tidak ditemukan")
	}
	if existing.OwnerID != userID {
		return nil, errors.New("Anda tidak memiliki izin untuk mengubah titik ini")
	}

	status := req.Status
	if status == "" {
		status = existing.Status
	}

	point := &MapPoint{
		CategoryID:   &req.CategoryID,
		Name:         req.Name,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		Address:      req.Address,
		TahunBerdiri: &req.TahunBerdiri,
		Description:  req.Description,
		CoverImage:   &req.CoverImage,
		Status:       status,
	}

	if err := s.repo.Update(ctx, id, point); err != nil {
		return nil, err
	}
	return point, nil
}

func (s *Service) DeletePoint(ctx context.Context, userID string, id int) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil || existing.OwnerID != userID {
		return errors.New("Titik tidak ditemukan atau Anda tidak memiliki izin")
	}
	return s.repo.Delete(ctx, id)
}

// Blog Services
func (s *Service) UpsertBlog(ctx context.Context, pointID int, req UpsertBlogReq) (*Blog, error) {
	blog := &Blog{
		MapPointID: pointID,
		Title:      req.Title,
		Content:    req.Content,
		CoverPhoto: req.CoverPhoto,
	}
	if err := s.repo.UpsertBlog(ctx, blog); err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *Service) GetBlog(ctx context.Context, pointID int) (*Blog, error) {
	return s.repo.GetBlogByPointID(ctx, pointID)
}

func (s *Service) Verify(ctx context.Context, id int, status string, rejectionNote string) error {
	var note *string
	if status == "rejected" && rejectionNote != "" {
		note = &rejectionNote
	}
	return s.repo.Verify(ctx, id, status, note)
}
func (s *Service) GetBlogDetail(ctx context.Context, pointID int) (*BlogDetailResp, error) {
	blog, err := s.repo.GetBlogByPointID(ctx, pointID)
	if err != nil {
		return nil, errors.New("Blog tidak ditemukan")
	}

	point, err := s.repo.GetByID(ctx, pointID)
	if err != nil {
		return nil, errors.New("Titik peta tidak ditemukan")
	}

	return &BlogDetailResp{
		Blog:  blog,
		Point: point,
	}, nil
}
