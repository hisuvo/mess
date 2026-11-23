package roles

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
    Create(ctx context.Context, role *Role) error
    GetByID(ctx context.Context, coachingID, id uint) (*Role, error)
    GetAll(ctx context.Context, coachingID uint) ([]Role, error)
    Update(ctx context.Context, role *Role) error
    Delete(ctx context.Context, coachingID, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, rule *Role) error {
	return nil
}

func (r *repository) GetByID(ctx context.Context, coachingID, id uint) (*Role, error){
	return nil, nil
}

func (r *repository) GetAll(ctx context.Context, coachingID uint) ([]Role, error) {
	return nil, nil
}

func (r *repository) Update(ctx context.Context, role *Role) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, coachingID, id uint) error {
	return nil
}