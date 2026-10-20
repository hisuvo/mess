package meals

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GenerateMill(ctx context.Context, meal *Meal) error
	FindById(ctx context.Context, id *uint) (*Meal, error)
	GetAll(ctx context.Context) (*Meal, error)
	Update(ctx context.Context, id uint, meal *Meal) (*Meal, error)
	Delete(ctx context.Context, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GenerateMill(ctx context.Context, meal *Meal) error {
	return nil
}

func (r *repository) FindById(ctx context.Context, id *uint) (*Meal, error) {
	return nil, nil
}

func (r *repository) GetAll(ctx context.Context) (*Meal, error) {
	return nil, nil
}

func (r *repository) Update(ctx context.Context, id uint, meal *Meal) (*Meal, error) {
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return nil
}
