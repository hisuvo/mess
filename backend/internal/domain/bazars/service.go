package bazars

import (
	"context"
	"messify/internal/domain/bazars/dto"
)

type Service interface {
	BazarCreate(ctx context.Context, req *dto.RequestCreateBazar) error
	GetAll(ctx context.Context) (*dto.ResponseBazar, error)
	FindById(ctx context.Context, id *uint) (*dto.ResponseBazar, error)
	BazarUpdate(ctx context.Context, req *dto.RequestUpdateBazar) (*dto.ResponseBazar, error)
	BazarDelete(ctx context.Context, id *uint) error
}

type serviec struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviec{
		repository: repository,
	}
}

func(s *serviec) BazarCreate(ctx context.Context, req *dto.RequestCreateBazar) error{
	return nil
}

func(s *serviec) GetAll(ctx context.Context) (*dto.ResponseBazar, error){
	return nil, nil
}

func(s *serviec) FindById(ctx context.Context, id *uint) (*dto.ResponseBazar, error){
	return nil, nil
}

func(s *serviec) BazarUpdate(ctx context.Context, req *dto.RequestUpdateBazar) (*dto.ResponseBazar, error){
	return nil, nil
}

func(s *serviec) BazarDelete(ctx context.Context, id *uint) error{
	return nil
}
