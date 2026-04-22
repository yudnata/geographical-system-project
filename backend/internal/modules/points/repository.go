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
	query := `
		SELECT gp.id, gp.type_id, gp.name, gp.latitude, gp.longitude, gp.address, gp.owner_id, COALESCE(u.name, 'Sistem'),
		gp.tahun_berdiri, gp.status_kepemilikan, gp.description, gp.is_active, gp.created_at, gp.updated_at
		FROM geo_points gp
		LEFT JOIN users u ON gp.owner_id = u.id
		WHERE gp.is_active = true
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []GeoPoint{}
	for rows.Next() {
		var p GeoPoint
		err := rows.Scan(
			&p.ID, &p.TypeID, &p.Name, &p.Latitude, &p.Longitude,
			&p.Address, &p.OwnerID, &p.OwnerName, &p.TahunBerdiri, &p.StatusKepemilikan,
			&p.Description, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

func (r *Repository) GetMyPoints(ctx context.Context, ownerID string) ([]GeoPoint, error) {
	query := `
		SELECT gp.id, gp.type_id, gp.name, gp.latitude, gp.longitude, gp.address, gp.owner_id, COALESCE(u.name, 'Sistem'),
		gp.tahun_berdiri, gp.status_kepemilikan, gp.description, gp.is_active, gp.created_at, gp.updated_at
		FROM geo_points gp
		LEFT JOIN users u ON gp.owner_id = u.id
		WHERE gp.is_active = true AND gp.owner_id = $1
	`
	rows, err := r.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []GeoPoint{}
	for rows.Next() {
		var p GeoPoint
		err := rows.Scan(
			&p.ID, &p.TypeID, &p.Name, &p.Latitude, &p.Longitude,
			&p.Address, &p.OwnerID, &p.OwnerName, &p.TahunBerdiri, &p.StatusKepemilikan,
			&p.Description, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}
func (r *Repository) GetByID(ctx context.Context, id int) (*GeoPoint, error) {
	var p GeoPoint
	err := r.db.QueryRow(ctx, 
		`SELECT id, type_id, name, latitude, longitude, address, owner_id, tahun_berdiri, status_kepemilikan, description, is_active, created_at, updated_at 
		 FROM geo_points WHERE id=$1`, id).Scan(
		&p.ID, &p.TypeID, &p.Name, &p.Latitude, &p.Longitude,
		&p.Address, &p.OwnerID, &p.TahunBerdiri, &p.StatusKepemilikan,
		&p.Description, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(ctx context.Context, id int, p *GeoPoint) error {
	_, err := r.db.Exec(ctx,
		`UPDATE geo_points
		SET type_id=$1, name=$2, latitude=$3, longitude=$4, address=$5, tahun_berdiri=$6, status_kepemilikan=$7, description=$8, updated_at=$9
		WHERE id=$10`,
		p.TypeID, p.Name, p.Latitude, p.Longitude, p.Address, p.TahunBerdiri, p.StatusKepemilikan, p.Description, time.Now(), id)
	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM geo_points WHERE id=$1`, id)
	return err
}
