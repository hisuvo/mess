package tenants

import (
	"context"
	"errors"
	"messify/internal/domain/tenants/dto"
	"messify/internal/pkg/querybuilder"
	"net/url"
	"strings"
)

var (
	ErrTenantEmailRequired = errors.New("email is required")
	ErrTenantNameRequired  = errors.New("tenant name is required")
)

type Service interface {
	CreateTenant(ctx context.Context, req *dto.RequestCreateTenant)  error
	GetAll(ctx context.Context, params url.Values) ([]*dto.ResponseTenant, querybuilder.Pagination, error)
	GetById(ctx context.Context, id uint) (*dto.ResponseTenant, error)
	FindByEmail(ctx context.Context, email string) (*dto.ResponseTenant, error)
	Update(ctx context.Context, id uint, req *dto.RequestUpdateTenant) (*dto.ResponseTenant, error)
	SoftDelete(ctx context.Context, id uint) (*dto.ResponseTenant, error)
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

func (s *service) GetAll(ctx context.Context, params url.Values) ([]*dto.ResponseTenant, querybuilder.Pagination, error){
	// Delegate database operations to the repository.
	tenants, pagination, err := s.repository.GetAll(ctx, params)

	if err != nil {
		return nil, querybuilder.Pagination{}, err
	}

	reponse := ToTenantsResponse(tenants)

	return reponse, pagination, nil
}

func (s *service) GetById(ctx context.Context, id uint) (*dto.ResponseTenant, error){
	tenant, err := s.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	response := ToTenantResponse(tenant)

	return response, nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*dto.ResponseTenant, error){
	return nil, nil
}

func (s *service) Update(ctx context.Context, id uint, req *dto.RequestUpdateTenant ) (*dto.ResponseTenant, error) {
	// Validate the tenant ID.
	if id == 0 {
		return nil, ErrTenantNotFound
	}

	// Reject a nil request.
	if req == nil {
		return nil, errors.New("update request cannot be nil")
	}

	// Build the tenant model from the request.
	// Set the existing tenant ID.
	// This ensures that the correct record is updated.
	tenant := &Tenant{
		ID: id,
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Address: req.Address,
	}

	// Delegate the database operation to the repository.
	updatedTenant, err := s.repository.Update(ctx, id, tenant)
	if err != nil {
		return nil, err
	}

	// Return the updated tenant.
	return ToTenantResponse(updatedTenant), nil
}

func (s *service) SoftDelete(ctx context.Context, id uint) (*dto.ResponseTenant, error){
	tenant, err := s.repository.SoftDelete(ctx, id)

	if err != nil {
		return nil, err
	}

	return ToTenantResponse(tenant), nil
}
