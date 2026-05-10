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

func (r *Repository) Create(ctx context.Context, p *MapPoint) error {
	query := `
		INSERT INTO map_points (
			category_id, name, latitude, longitude, address, owner_id,
			tahun_berdiri, status_kepemilikan, description, is_active, status, rejection_note,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id
	`

	now := time.Now()
	err := r.db.QueryRow(ctx, query,
		p.CategoryID,
		p.Name,
		p.Latitude,
		p.Longitude,
		p.Address,
		p.OwnerID,
		p.TahunBerdiri,
		p.StatusKepemilikan,
		p.Description,
		p.IsActive,
		p.Status,
		p.RejectionNote,
		now,
		now,
	).Scan(&p.ID)

	if err == nil {
		p.CreatedAt = now
		p.UpdatedAt = now
	}
	return err
}

func (r *Repository) GetAll(ctx context.Context) ([]MapPoint, error) {
	query := `
		SELECT gp.id, gp.category_id, gp.name, gp.latitude, gp.longitude, gp.address, gp.owner_id, COALESCE(u.name, 'Sistem'),
		gp.tahun_berdiri, gp.status_kepemilikan, gp.description, gp.is_active, gp.status, gp.rejection_note, gp.created_at, gp.updated_at
		FROM map_points gp
		LEFT JOIN users u ON gp.owner_id = u.id
		WHERE gp.is_active = true
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []MapPoint{}
	for rows.Next() {
		var p MapPoint
		err := rows.Scan(
			&p.ID, &p.CategoryID, &p.Name, &p.Latitude, &p.Longitude,
			&p.Address, &p.OwnerID, &p.OwnerName, &p.TahunBerdiri, &p.StatusKepemilikan,
			&p.Description, &p.IsActive, &p.Status, &p.RejectionNote, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

func (r *Repository) GetMyPoints(ctx context.Context, ownerID string) ([]MapPoint, error) {
	query := `
		SELECT gp.id, gp.category_id, gp.name, gp.latitude, gp.longitude, gp.address, gp.owner_id, COALESCE(u.name, 'Sistem'),
		gp.tahun_berdiri, gp.status_kepemilikan, gp.description, gp.is_active, gp.status, gp.rejection_note, gp.created_at, gp.updated_at
		FROM map_points gp
		LEFT JOIN users u ON gp.owner_id = u.id
		WHERE gp.is_active = true AND gp.owner_id = $1
	`
	rows, err := r.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []MapPoint{}
	for rows.Next() {
		var p MapPoint
		err := rows.Scan(
			&p.ID, &p.CategoryID, &p.Name, &p.Latitude, &p.Longitude,
			&p.Address, &p.OwnerID, &p.OwnerName, &p.TahunBerdiri, &p.StatusKepemilikan,
			&p.Description, &p.IsActive, &p.Status, &p.RejectionNote, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*MapPoint, error) {
	var p MapPoint
	err := r.db.QueryRow(ctx, 
		`SELECT id, category_id, name, latitude, longitude, address, owner_id, tahun_berdiri, status_kepemilikan, description, is_active, status, rejection_note, created_at, updated_at 
		 FROM map_points WHERE id=$1`, id).Scan(
		&p.ID, &p.CategoryID, &p.Name, &p.Latitude, &p.Longitude,
		&p.Address, &p.OwnerID, &p.TahunBerdiri, &p.StatusKepemilikan,
		&p.Description, &p.IsActive, &p.Status, &p.RejectionNote, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(ctx context.Context, id int, p *MapPoint) error {
	_, err := r.db.Exec(ctx,
		`UPDATE map_points
		SET category_id=$1, name=$2, latitude=$3, longitude=$4, address=$5, tahun_berdiri=$6, status_kepemilikan=$7, description=$8, is_active=$9, status=$10, rejection_note=$11, updated_at=$12
		WHERE id=$13`,
		p.CategoryID, p.Name, p.Latitude, p.Longitude, p.Address, p.TahunBerdiri, p.StatusKepemilikan, p.Description, p.IsActive, p.Status, p.RejectionNote, time.Now(), id)
	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM map_points WHERE id=$1`, id)
	return err
}
