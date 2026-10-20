package tenants

import (
	"context"
	"messify/internal/domain/tenants/dto"
)

type Service interface {
	CreateTenant(ctx context.Context, req *dto.RequestCreateTenant) (*dto.ResponseTenant, error)
	GetAll(ctx context.Context) ([]*dto.ResponseGetTenants, error)
	GetById(ctx context.Context, id uint) (*dto.ResponseTenant, error)
	FindByEmail(ctx context.Context, email string) (*dto.ResponseTenant, error)
	Update(ctx context.Context, tenant *dto.RequestUpdateTenant) (*dto.ResponseTenant, error)
	Delete(ctx context.Context, id uint) (*dto.ResponseTenant, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateTenant(ctx context.Context,req *dto.RequestCreateTenant) (*dto.ResponseTenant, error){
	return nil, nil
}

func (s *service) GetAll(ctx context.Context) ([]*dto.ResponseGetTenants, error){
	return nil, nil
}

func (s *service) GetById(ctx context.Context, id uint) (*dto.ResponseTenant, error){
	return nil, nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*dto.ResponseTenant, error){
	return nil, nil
}

func (s *service) Update(ctx context.Context, tenant *dto.RequestUpdateTenant) (*dto.ResponseTenant, error){
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id uint) (*dto.ResponseTenant, error){
	return nil, nil
}
