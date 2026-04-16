package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *User) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (id, email, name, password, sso_provider, sso_id, avatar_url, phone, institution, is_profile_completed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		u.ID, u.Email, u.Name, nullIfEmpty(u.Password), u.SSOProvider, u.SSOID, u.AvatarURL, nullIfEmptyPtr(u.Phone), nullIfEmptyPtr(u.Institution), u.IsProfileCompleted,
	)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	var pw, ssoProv, ssoId, avatarUrl, phone, institution *string
	err := r.db.QueryRow(ctx, `SELECT id, email, name, password, sso_provider, sso_id, avatar_url, phone, institution, is_profile_completed, created_at FROM users WHERE email = $1`, email).
		Scan(&u.ID, &u.Email, &u.Name, &pw, &ssoProv, &ssoId, &avatarUrl, &phone, &institution, &u.IsProfileCompleted, &u.CreatedAt)

	if pw != nil {
		u.Password = *pw
	}
	u.SSOProvider = ssoProv
	u.SSOID = ssoId
	u.AvatarURL = avatarUrl
	u.Phone = phone
	u.Institution = institution

	return &u, err
}

func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	var u User
	var pw, ssoProv, ssoId, avatarUrl, phone, institution *string
	err := r.db.QueryRow(ctx, `SELECT id, email, name, password, sso_provider, sso_id, avatar_url, phone, institution, is_profile_completed, created_at FROM users WHERE id = $1`, id).
		Scan(&u.ID, &u.Email, &u.Name, &pw, &ssoProv, &ssoId, &avatarUrl, &phone, &institution, &u.IsProfileCompleted, &u.CreatedAt)

	if pw != nil {
		u.Password = *pw
	}
	u.SSOProvider = ssoProv
	u.SSOID = ssoId
	u.AvatarURL = avatarUrl
	u.Phone = phone
	u.Institution = institution

	return &u, err
}

func (r *Repository) UpdateSSO(ctx context.Context, id, provider, ssoId, avatarUrl string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET sso_provider=$1, sso_id=$2, avatar_url=$3 WHERE id=$4`, provider, ssoId, avatarUrl, id)
	return err
}

func (r *Repository) UpdateProfile(ctx context.Context, id, name, phone, institution string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET name=$1, phone=$2, institution=$3, is_profile_completed=true WHERE id=$4`, name, phone, institution, id)
	return err
}

func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func nullIfEmptyPtr(s *string) *string {
	if s == nil || *s == "" {
		return nil
	}
	return s
}
