package deposits

import (
	"context"
	"messify/internal/domain/deposits/dto"
)

type Service interface {
	GenerateDeposit(ctx context.Context, deposit *dto.RequestCreateDeposit) error
	GetAll(ctx context.Context) (*dto.DepositsResponse, error)
	FindById(ctx context.Context, id *uint) (*dto.DepositsResponse, error)
	Update(ctx context.Context, id *uint, deposit *dto.RequestUpdateDeposit) (*dto.DepositsResponse, error)
	Delete(crx context.Context, id *uint) error
}

type service struct {
	repository Repository
}

func NewService (repository Repository) Service {
	return &service{
		repository: repository,
	}
}


func (s *service) GenerateDeposit(ctx context.Context, deposit *dto.RequestCreateDeposit) error {
	return nil
}

func (s *service) GetAll(ctx context.Context) (*dto.DepositsResponse, error) {
	return nil, nil
}

func (s *service) FindById(ctx context.Context, id *uint) (*dto.DepositsResponse, error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, id *uint, deposit *dto.RequestUpdateDeposit) (*dto.DepositsResponse, error) {
	return nil, nil
}

func (s *service) Delete(crx context.Context, id *uint) error {
	return nil
}
