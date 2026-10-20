package tenants

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTenant(ctx context.Context, tenants *Tenant) error
	GetAll(ctx context.Context) ([]*Tenant, error)
	GetById(ctx context.Context, id uint) (*Tenant, error)
	FindByEmail(ctx context.Context, email string) (*Tenant, error)
	Update(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Delete(ctx context.Context, id uint) (*Tenant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTenant(ctx context.Context, tenants *Tenant) error {
	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]*Tenant, error){
	return nil, nil
}

func (r *repository) GetById(ctx context.Context, id uint) (*Tenant, error){
	return nil, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*Tenant, error){
	return nil, nil
}

func (r *repository) Update(ctx context.Context, tenant *Tenant) (*Tenant, error){
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id uint) (*Tenant, error){
	return nil, nil
}