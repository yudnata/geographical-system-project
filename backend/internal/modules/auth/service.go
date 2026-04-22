package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
}

func NewService(repo *Repository, cfg *config.Config) *Service {
	return &Service{repo: repo, cfg: cfg}
}

func (s *Service) generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *Service) Register(input RegisterReq) (*User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	user := &User{
		ID:        uuid.New().String(),
		Email:     input.Email,
		Name:      input.Name,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}
	return user, s.repo.Create(context.Background(), user)
}

func (s *Service) Login(input LoginReq) (string, *User, error) {
	user, err := s.repo.FindByEmail(context.Background(), input.Email)
	if err != nil {
		return "", nil, errors.New("Email atau password salah")
	}
	if user.Password == "" {
		provider := "Google"
		if user.SSOProvider != nil && *user.SSOProvider != "" {
			provider = *user.SSOProvider
		}
		return "", nil, fmt.Errorf("Akun ini terdaftar via %s. Silakan login menggunakan tombol '%s'", provider, provider)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return "", nil, errors.New("Password yang Anda masukkan salah")
	}
	token, err := s.generateToken(user.ID)
	s.prepareUser(user)
	return token, user, err
}

func (s *Service) SSOLogin(input SSOLoginReq) (string, *User, error) {
	user, err := s.repo.FindByEmail(context.Background(), input.Email)
	if err != nil {
		user = &User{
			ID:                 uuid.New().String(),
			Email:              input.Email,
			Name:               input.Name,
			SSOProvider:        &input.SSOProvider,
			SSOID:              &input.SSOID,
			AvatarURL:          &input.AvatarURL,
			IsProfileCompleted: false,
			CreatedAt:          time.Now(),
		}
		if createErr := s.repo.Create(context.Background(), user); createErr != nil {
			return "", nil, createErr
		}
	} else {
		s.repo.UpdateSSO(context.Background(), user.ID, input.SSOProvider, input.SSOID, input.AvatarURL)
		user.AvatarURL = &input.AvatarURL
	}

	token, err := s.generateToken(user.ID)
	s.prepareUser(user)
	return token, user, err
}

func (s *Service) UpdateProfile(userID string, input UpdateProfileReq) (*User, error) {
	if input.Name == "" || input.Phone == "" {
		return nil, errors.New("Nama dan nomor telp wajib diisi")
	}

	err := s.repo.UpdateProfile(context.Background(), userID, input.Name, input.Phone, input.Institution)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.FindByID(context.Background(), userID)
	if err == nil {
		s.prepareUser(user)
	}
	return user, err
}

func (s *Service) GetMe(userID string) (*User, error) {
	user, err := s.repo.FindByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	s.prepareUser(user)
	return user, nil
}

func (s *Service) UpdatePassword(userID string, req UpdatePasswordReq) error {
	user, err := s.repo.FindByID(context.Background(), userID)
	if err != nil {
		return errors.New("User tidak ditemukan")
	}

	if user.Password != "" {
		if req.OldPassword == nil || *req.OldPassword == "" {
			return errors.New("Password lama wajib diisi")
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*req.OldPassword)) != nil {
			return errors.New("Password lama salah")
		}
	}

	if len(req.NewPassword) < 6 {
		return errors.New("Password baru minimal 6 karakter")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		return errors.New("Gagal mengenkripsi password")
	}

	return s.repo.UpdatePassword(context.Background(), userID, string(hash))
}

func (s *Service) prepareUser(u *User) {
	if u.AvatarURL == nil || *u.AvatarURL == "" {
		u.AvatarURL = &s.cfg.DefaultAvatarURL
	}
	u.HasPassword = u.Password != ""
}
