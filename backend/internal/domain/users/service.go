package users

import (
	"context"
	"messify/internal/domain/users/dto"
	"messify/internal/pkg/security"
)

type Service interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*dto.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	FindAll(ctx context.Context, req *dto.ListUserRequest) ([]dto.UserResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	//hash password here
	hashPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	//create user
	user := &User{
		Name: req.Name,
		Email: req.Email,
		Password: hashPassword,
		Role: UserRole(req.Role),
		Status: UserStatusActive,
		ImageURL: req.ImageUrl,
		Phone: req.Phone,
		Address: req.Address,
	}

	result, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return ToUserResponse(result), nil
}
func (s *service) Update(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	return nil, nil
}
func (s *service) Delete(ctx context.Context, id uint) error {
	return nil
}
func (s *service) FindByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	return nil, nil
}
func (s *service) FindByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	return nil, nil
}
func (s *service) FindAll(ctx context.Context, req *dto.ListUserRequest) ([]dto.UserResponse, error) {
	return nil, nil
}
