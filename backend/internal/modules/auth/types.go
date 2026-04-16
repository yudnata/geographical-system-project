package auth

import "time"

type User struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	Name               string    `json:"name"`
	Password           string    `json:"-"`
	SSOProvider        *string   `json:"sso_provider"`
	SSOID              *string   `json:"sso_id"`
	AvatarURL          *string   `json:"avatar_url"`
	Phone              *string   `json:"phone"`
	Institution        *string   `json:"institution"`
	IsProfileCompleted bool      `json:"is_profile_completed"`
	CreatedAt          time.Time `json:"created_at"`
}

type UpdateProfileReq struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Institution string `json:"institution"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SSOLoginReq struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	SSOProvider string `json:"sso_provider"`
	SSOID       string `json:"sso_id"`
	AvatarURL   string `json:"avatar_url"`
}
