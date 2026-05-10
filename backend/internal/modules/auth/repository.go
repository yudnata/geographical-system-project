package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

const userSelectQuery = `
	SELECT id, email, full_name, password, sso_provider, sso_id, avatar_url, phone, is_profile_completed, role, created_at 
	FROM users
`

func (r *Repository) scanUser(row pgx.Row) (*User, error) {

	var u User
	var pw, ssoProv, ssoId, avatarUrl, phone *string
	err := row.Scan(&u.ID, &u.Email, &u.FullName, &pw, &ssoProv, &ssoId, &avatarUrl, &phone, &u.IsProfileCompleted, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	if pw != nil {
		u.Password = *pw
		u.HasPassword = true
	}
	u.SSOProvider = ssoProv
	u.SSOID = ssoId
	u.AvatarURL = avatarUrl
	u.Phone = phone

	return &u, nil
}

func (r *Repository) Create(ctx context.Context, u *User) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (id, email, full_name, password, sso_provider, sso_id, avatar_url, phone, is_profile_completed, role) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		u.ID, u.Email, u.FullName, nullIfEmpty(u.Password), u.SSOProvider, u.SSOID, u.AvatarURL, nullIfEmptyPtr(u.Phone), u.IsProfileCompleted, u.Role,
	)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRow(ctx, userSelectQuery+" WHERE email = $1", email)
	return r.scanUser(row)
}

func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	row := r.db.QueryRow(ctx, userSelectQuery+" WHERE id = $1", id)
	return r.scanUser(row)
}

func (r *Repository) FindAllUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, email, full_name, sso_provider, avatar_url, phone, is_profile_completed, role, created_at FROM users ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		var ssoProv, avatarUrl, phone *string
		if err := rows.Scan(&u.ID, &u.Email, &u.FullName, &ssoProv, &avatarUrl, &phone, &u.IsProfileCompleted, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		u.SSOProvider = ssoProv
		u.AvatarURL = avatarUrl
		u.Phone = phone
		users = append(users, u)
	}
	return users, nil
}

func (r *Repository) UpdateSSO(ctx context.Context, id, provider, ssoId, avatarUrl string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET sso_provider=$1, sso_id=$2, avatar_url=$3 WHERE id=$4`, provider, ssoId, avatarUrl, id)
	return err
}

func (r *Repository) UpdateProfile(ctx context.Context, id, fullName, phone string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET full_name=$1, phone=$2, is_profile_completed=true WHERE id=$3`, fullName, phone, id)
	return err
}

func (r *Repository) UpdatePassword(ctx context.Context, id, passwordHash string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET password=$1 WHERE id=$2`, passwordHash, id)
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
