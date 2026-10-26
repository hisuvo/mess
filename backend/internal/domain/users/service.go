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
	GetAll(ctx context.Context) ([]*dto.UserResponse, error)
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
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func (s *service) GetAll(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return ToUserResponses(users), nil
}
