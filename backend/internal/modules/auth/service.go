package auth

import (
	"context"
	"errors"
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
	if err != nil || user.Password == "" || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return "", nil, errors.New("Email atau password salah")
	}
	token, err := s.generateToken(user.ID)
	s.ensureAvatar(user)
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
	} else if user.SSOID == nil {
		s.repo.UpdateSSO(context.Background(), user.ID, input.SSOProvider, input.SSOID, input.AvatarURL)
	}

	token, err := s.generateToken(user.ID)
	s.ensureAvatar(user)
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
		s.ensureAvatar(user)
	}
	return user, err
}

func (s *Service) ensureAvatar(u *User) {
	if u.AvatarURL == nil || *u.AvatarURL == "" {
		u.AvatarURL = &s.cfg.DefaultAvatarURL
	}
}
