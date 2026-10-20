package notices

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GenerateNotice(ctx context.Context, notice *Notice) error
	GetAll(ctx context.Context) (*Notice, error)
	FindById(ctx context.Context, id *uint) (*Notice, error)
	Update(ctx context.Context, id *uint, notice *Notice) (*Notice, error)
	Delete(ctx context.Context, id *uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repository{
		db: db,
	}
}


func (r *repository) GenerateNotice(ctx context.Context, notice *Notice) error{
	return nil
}

func (r *repository) GetAll(ctx context.Context) (*Notice, error){
	return nil, nil
}

func (r *repository) FindById(ctx context.Context, id *uint) (*Notice, error){
	return nil, nil
}

func (r *repository) Update(ctx context.Context, id *uint, notice *Notice) (*Notice, error){
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id *uint) error{
	return nil
}
