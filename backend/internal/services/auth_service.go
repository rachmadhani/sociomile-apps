package services

import (
	"errors"
	authDTO "sociomile-apps/internal/dto/auth"
	"sociomile-apps/internal/middleware"
	model "sociomile-apps/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(input authDTO.RegisterRequest) (*authDTO.AuthResponse, error) {
	var existingUsers model.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUsers).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	tenant := model.Tenant{
		Name: input.TenantName,
	}

	if err := s.db.Create(&tenant).Error; err != nil {
		return nil, errors.New("failed to create tenant")
	}

	user := model.User{
		TenantID:     tenant.ID,
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         model.RoleAdmin,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	token, err := middleware.GenerateToken(&user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &authDTO.AuthResponse{
		AccessToken: token,
		User:        &user,
	}, nil
}

func (s *AuthService) Login(input authDTO.LoginRequest) (*authDTO.AuthResponse, error) {
	var user model.User
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := middleware.GenerateToken(&user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	session := model.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	if err := s.db.Create(&session).Error; err != nil {
		return nil, errors.New("failed to create session")
	}

	return &authDTO.AuthResponse{
		AccessToken: token,
		User:        &user,
	}, nil
}

func (s *AuthService) Logout(token string) error {
	var session model.Session
	if err := s.db.Where("token = ?", token).First(&session).Error; err != nil {
		return errors.New("session not found")
	}

	if err := s.db.Delete(&session).Error; err != nil {
		return errors.New("failed to delete session")
	}

	return nil
}
