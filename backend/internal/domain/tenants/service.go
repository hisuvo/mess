package tenants

import (
	"context"
	"errors"
	"messify/internal/domain/tenants/dto"
	"strings"
)

var (
	ErrTenantEmailRequired = errors.New("email is required")
	ErrTenantNameRequired  = errors.New("tenant name is required")
)

type Service interface {
	CreateTenant(ctx context.Context, req *dto.RequestCreateTenant)  error
	GetAll(ctx context.Context) ([]*dto.ResponseTenant, error)
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


func (s *service) CreateTenant(ctx context.Context,req *dto.RequestCreateTenant) error {
	// Validate the request.
	if req == nil {
		return errors.New("tenant request is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return ErrTenantEmailRequired
	}

	if strings.TrimSpace(req.Name) == "" {
		return ErrTenantNameRequired
	}

	// Normalize the email to avoid whitespace issues.
	email := strings.TrimSpace(req.Email)

	// Map the request DTO to the domain model.
	tenant := &Tenant{
		Name:    strings.TrimSpace(req.Name),
		Email:   email,
		Address: strings.TrimSpace(req.Address),
		Phone:   strings.TrimSpace(req.Phone),
	}

	// Persist the tenant and propagate repository errors.
	if err := s.repository.CreateTenant(ctx, tenant); err != nil {
		return err
	}

		return nil
}

func (s *service) GetAll(ctx context.Context) ([]*dto.ResponseTenant, error){
	tenants, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	reponse := ToTenantsResponse(tenants)

	return reponse, nil
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
