package deposits

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GenerateDeposit(ctx context.Context, deposit *Deposit) error
	GetAll(ctx context.Context) (*Deposit, error)
	FindById(ctx context.Context, id *uint) (*Deposit, error)
	Update(ctx context.Context, id *uint, deposit *Deposit) (*Deposit, error)
	Delete(crx context.Context, id *uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}


func (r *repository) GenerateDeposit(ctx context.Context, deposit *Deposit) error {
	return nil
}
func (r *repository) GetAll(ctx context.Context) (*Deposit, error) {
	return nil, nil
}
func (r *repository) FindById(ctx context.Context, id *uint) (*Deposit, error) {
	return nil, nil
}
func (r *repository) Update(ctx context.Context, id *uint, deposit *Deposit) (*Deposit, error) {
	return nil, nil
}
func (r *repository) Delete(crx context.Context, id *uint) error {
	return nil
}