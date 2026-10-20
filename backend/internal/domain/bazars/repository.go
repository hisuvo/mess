package bazars

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	BazarCreate(ctx context.Context, bazar *Bazar) error
	GetAll(ctx context.Context) (*Bazar, error)
	FindById(ctx context.Context, id *uint) (*Bazar, error)
	BazarUpdate(ctx context.Context, bazar *Bazar) (*Bazar, error)
	BazarDelete(ctx context.Context, id *uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) BazarCreate(ctx context.Context, bazar *Bazar) error{
	return nil
}

func(r *repository) GetAll(ctx context.Context) (*Bazar, error){
	return nil, nil
}

func(r *repository) FindById(ctx context.Context, id *uint) (*Bazar, error){
	return nil, nil
}

func(r *repository) BazarUpdate(ctx context.Context, bazar *Bazar) (*Bazar, error){
	return nil, nil
}

func(r *repository) BazarDelete(ctx context.Context, id *uint) error{
	return nil
}