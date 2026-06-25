package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"notes_api/internal/dto"
	"notes_api/internal/model"
	"notes_api/internal/repository"
	"notes_api/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {

	existingUser, err := s.userRepo.GetUserByEmail(req.Email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}
