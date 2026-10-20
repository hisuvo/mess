package meals

import (
	"context"
	"messify/internal/domain/meals/dto"
)

type Service interface {
	GenerateMill(ctx context.Context, meal *dto.RequestCreatemeal) error
	FindById(ctx context.Context, id *uint) (*dto.Responsemeal, error)
	GetAll(ctx context.Context) (*dto.Responsemeal, error)
	Update(ctx context.Context, id uint, meal *dto.RequestCreatemeal) (*dto.Responsemeal, error)
	Delete(ctx context.Context, id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GenerateMill(ctx context.Context, meal *dto.RequestCreatemeal) error{
	return nil
}

func (s *service) FindById(ctx context.Context, id *uint) (*dto.Responsemeal, error){
	return nil, nil
}

func (s *service) GetAll(ctx context.Context) (*dto.Responsemeal, error){
	return nil, nil
}

func (s *service) Update(ctx context.Context, id uint, meal *dto.RequestCreatemeal) (*dto.Responsemeal, error){
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id uint) error{
	return nil
}
