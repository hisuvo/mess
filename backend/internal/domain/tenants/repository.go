package tenants

import (
	"context"
	"errors"
	"fmt"
	"messify/internal/pkg/querybuilder"
	"net/url"

	"gorm.io/gorm"
)

var (
	ErrTenantEmailExists = errors.New("Email already used")
	ErrTenantIdRequired = errors.New("tenant id is required")
	ErrTenantNotFound = errors.New("your tenant not found")
)

type Repository interface {
	CreateTenant(ctx context.Context, tenant *Tenant) error
	GetAll(ctx context.Context, params url.Values) ([]*Tenant, querybuilder.Pagination, error)
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

func (r *repository) GetAll(ctx context.Context, params url.Values) ([]*Tenant, querybuilder.Pagination, error){
	var tenants []*Tenant
	var total int64


	// Only these columns can be sorted.
	allowedSortFields := []string {
		"id",
		"name",
		"email",
		"created_at",
	}

	// Parse and validate query parameters.
	options := querybuilder.Parse(params, allowedSortFields)

	// Start with a context-aware database query
	db := r.db.WithContext(ctx).Model(&Tenant{})

	// Apply search before counting and pagination.
	db = querybuilder.ApplySearch(
		db,
		options.Search,
		[]string{"name", "email"},
	)

	// Count matching records before applying LIMIT/OFFSET.
	if err := db.Count(&total).Error; err != nil {
		return nil, querybuilder.Pagination{}, err
	}


	// Apply sorting and pagination.
	db = querybuilder.ApplySorting(
		db,
		options,
		"created_at",
		allowedSortFields,
	)

	db = querybuilder.ApplyPagination(db, options)

	if err := db.Find(&tenants).Error; err != nil {
		return nil, querybuilder.Pagination{}, err
	}

	meta := querybuilder.BuildPagination(
		options.Page,
		options.Limit,
		total,
	)

	fmt.Println("this is meta data of tenants ->", meta)

	return tenants, meta, nil
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