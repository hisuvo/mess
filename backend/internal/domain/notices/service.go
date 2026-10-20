package notices

import (
	"context"
	"messify/internal/domain/notices/dto"
)

type Service interface {
	GenerateNotice(ctx context.Context, req *dto.RequestCreateNotice) error
	GetAll(ctx context.Context) (*dto.ResponseNotice, error)
	FindById(ctx context.Context, id *uint) (*dto.ResponseNotice, error)
	Update(ctx context.Context, id *uint, data *dto.RequestUpdateNotice) (*dto.ResponseNotice, error)
	Delete(ctx context.Context, id *uint) error
}

type service struct {
	repository Repository
}

func NewService (repository Repository) Service {
	return &service{
		repository: repository,
	}
}


func (s *service) GenerateNotice(ctx context.Context, req *dto.RequestCreateNotice) error{
	return nil
}

func (s *service) GetAll(ctx context.Context) (*dto.ResponseNotice, error){
	return nil, nil
}

func (s *service) FindById(ctx context.Context, id *uint) (*dto.ResponseNotice, error){
	return nil, nil
}

func (s *service) Update(ctx context.Context, id *uint, data *dto.RequestUpdateNotice) (*dto.ResponseNotice, error){
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id *uint) error{
	return nil
}
