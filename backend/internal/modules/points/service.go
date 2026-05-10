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

func (s *Service) CreatePoint(ctx context.Context, ownerID string, input CreatePointReq) (*MapPoint, error) {
	status := input.Status
	if status == "" {
		status = "draft"
	}
	
	point := &MapPoint{
		CategoryID:        &input.CategoryID,
		Name:              input.Name,
		Latitude:          input.Latitude,
		Longitude:         input.Longitude,
		Address:           input.Address,
		OwnerID:           ownerID,
		TahunBerdiri:      &input.TahunBerdiri,
		StatusKepemilikan: input.StatusKepemilikan,
		Description:       input.Description,
		IsActive:          true,
		Status:            status,
	}

	err := s.repo.Create(ctx, point)
	if err != nil {
		return nil, err
	}

	return point, nil
}

func (s *Service) GetAllPoints(ctx context.Context) ([]MapPoint, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetMyPoints(ctx context.Context, ownerID string) ([]MapPoint, error) {
	return s.repo.GetMyPoints(ctx, ownerID)
}

func (s *Service) UpdatePoint(ctx context.Context, userID string, id int, req UpdatePointReq) (*MapPoint, error) {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("Titik tidak ditemukan")
	}
	if existing.OwnerID != userID {
		return nil, errors.New("Anda tidak memiliki izin untuk mengubah titik ini")
	}

	categoryID := req.CategoryID
	tahunBerdiri := req.TahunBerdiri
	status := req.Status
	if status == "" {
		status = existing.Status
	}

	point := &MapPoint{
		CategoryID:        &categoryID,
		Name:              req.Name,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		Address:           req.Address,
		TahunBerdiri:      &tahunBerdiri,
		StatusKepemilikan: req.StatusKepemilikan,
		Description:       req.Description,
		IsActive:          req.IsActive,
		Status:            status,
	}

	if err := s.repo.Update(ctx, id, point); err != nil {
		return nil, err
	}

	return point, nil
}

func (s *Service) DeletePoint(ctx context.Context, userID string, id int) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("Titik tidak ditemukan")
	}
	if existing.OwnerID != userID {
		return errors.New("Anda tidak memiliki izin untuk menghapus titik ini")
	}
	return s.repo.Delete(ctx, id)
}
