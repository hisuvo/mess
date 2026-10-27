package tenants

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrTenantEmailExists = errors.New("Email already used")
	ErrTenantIdRequired = errors.New("tenant id is required")
	ErrTenantNotFound = errors.New("your tenant not found")
)

type Repository interface {
	CreateTenant(ctx context.Context, tenant *Tenant) error
	GetAll(ctx context.Context) ([]*Tenant, error)
	GetById(ctx context.Context, id uint) (*Tenant, error)
	FindByEmail(ctx context.Context, email string) (*Tenant, error)
	Update(ctx context.Context, id uint, tenant *Tenant) (*Tenant, error)
	SoftDelete(ctx context.Context, id uint) (*Tenant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTenant(ctx context.Context, tenant *Tenant) error {
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
	var tenants []*Tenant

	if err := r.db.WithContext(ctx).Find(&tenants).Error; err != nil {
		return nil, err
	}

	return tenants, nil
}

func (r *repository) GetById(ctx context.Context, id uint) (*Tenant, error) {
	var tenant Tenant

	result := r.db.WithContext(ctx).First(&tenant, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrTenantNotFound
		}

		return nil, result.Error
	}

	return &tenant, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*Tenant, error) {
	var tenant *Tenant

	result := r.db.WithContext(ctx).Where("email = ?", email).Find(&tenant)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrTenantNotFound
		}
	}

	return tenant, nil
}

func (r *repository) Update(ctx context.Context, id uint, tenant *Tenant) (*Tenant, error) {
	// Validate the tenant and its ID.
	if tenant == nil || tenant.ID == 0 {
		return nil, ErrTenantNotFound
	}

	// Check whether the tenant exists.
	var existing Tenant
	if err := r.db.WithContext(ctx).First(&existing, tenant.ID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTenantNotFound
		}

		return nil, err
	}

	// Update the selected fields, including empty strings if provided.
	if err := r.db.WithContext(ctx).
		Model(&existing).
		Select("Name", "Email", "Address", "Phone").
		Updates(tenant).Error; err != nil {
		return nil, err
	}

	// Return the updated tenant.
	if err := r.db.WithContext(ctx).
		First(&existing, tenant.ID).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *repository) SoftDelete(ctx context.Context, id uint ) (*Tenant, error) {

	// Validate the tenant ID before querying the database.
    if id == 0 {
        return nil, ErrTenantNotFound
    }

	var tenant Tenant

	// Find the tenant first and ensure it exists.
    if err := r.db.WithContext(ctx).First(&tenant, id).Error; err != nil {

        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrTenantNotFound
        }

        return nil, err
    }

	// Soft-delete the tenant using GORM's built-in mechanism.
    if err := r.db.WithContext(ctx).Delete(&tenant).Error; err != nil {
        return nil, err
    }

    // Return the tenant with its DeletedAt field updated.
    return &tenant, nil
}