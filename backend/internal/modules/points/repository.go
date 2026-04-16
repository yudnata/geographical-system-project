package points

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p *GeoPoint) error {
	query := `
		INSERT INTO geo_points (
			type_id, name, latitude, longitude, address, owner_id,
			tahun_berdiri, status_kepemilikan, description, is_active,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	now := time.Now()
	err := r.db.QueryRow(ctx, query,
		p.TypeID,
		p.Name,
		p.Latitude,
		p.Longitude,
		p.Address,
		p.OwnerID,
		p.TahunBerdiri,
		p.StatusKepemilikan,
		p.Description,
		p.IsActive,
		now,
		now,
	).Scan(&p.ID)

	if err == nil {
		p.CreatedAt = now
		p.UpdatedAt = now
	}
	return err
}

func (r *Repository) GetAll(ctx context.Context) ([]GeoPoint, error) {
	query := `SELECT id, type_id, name, latitude, longitude, address, owner_id, tahun_berdiri, status_kepemilikan, description, is_active, created_at, updated_at FROM geo_points WHERE is_active = true`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var points []GeoPoint
	for rows.Next() {
		var p GeoPoint
		err := rows.Scan(
			&p.ID, &p.TypeID, &p.Name, &p.Latitude, &p.Longitude,
			&p.Address, &p.OwnerID, &p.TahunBerdiri, &p.StatusKepemilikan,
			&p.Description, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}
