package auth

import "time"

type User struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	FullName           string    `json:"full_name"`
	Password           string    `json:"-"`
	SSOProvider        *string   `json:"sso_provider"`
	SSOID              *string   `json:"sso_id"`
	AvatarURL          *string   `json:"avatar_url"`
	Phone              *string   `json:"phone"`
	IsProfileCompleted bool      `json:"is_profile_completed"`
	HasPassword        bool      `json:"has_password"`
	Role               string    `json:"role"`
	CreatedAt          time.Time `json:"created_at"`
}

type UpdateProfileReq struct {
	FullName  string `json:"full_name"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}


type RegisterReq struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SSOLoginReq struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	SSOProvider string `json:"sso_provider"`
	SSOID       string `json:"sso_id"`
	AvatarURL   string `json:"avatar_url"`
}


type UpdatePasswordReq struct {
	OldPassword *string `json:"old_password"`
	NewPassword string  `json:"new_password"`
}
