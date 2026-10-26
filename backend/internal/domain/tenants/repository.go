package tenants

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrTenantEmailExists = errors.New("Email already used")
)

type Repository interface {
	CreateTenant(ctx context.Context, tenant *Tenant) error
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


func (r *repository) CreateTenant(
	ctx context.Context,
	tenant *Tenant,
) error {
	err := r.db.WithContext(ctx).Create(tenant).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrTenantEmailExists
		}

		return err
	}

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