package points

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePoint(ctx context.Context, ownerID string, input CreatePointReq) (*GeoPoint, error) {
	point := &GeoPoint{
		TypeID:            &input.TypeID,
		Name:              input.Name,
		Latitude:          input.Latitude,
		Longitude:         input.Longitude,
		Address:           input.Address,
		OwnerID:           ownerID,
		TahunBerdiri:      &input.TahunBerdiri,
		StatusKepemilikan: input.StatusKepemilikan,
		Description:       input.Description,
		IsActive:          true,
	}

	err := s.repo.Create(ctx, point)
	if err != nil {
		return nil, err
	}

	return point, nil
}

func (s *Service) GetAllPoints(ctx context.Context) ([]GeoPoint, error) {
	return s.repo.GetAll(ctx)
}
