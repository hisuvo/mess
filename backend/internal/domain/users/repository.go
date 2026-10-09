package users

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, id uint, user *User) (*User, error)
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context, user *User) ([]*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, user *User) (*User, error) {
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (r *repository) Update(ctx context.Context, id uint, user *User) (*User, error) {
	return nil, nil
}
func (r *repository) Delete(ctx context.Context, id uint) error {
	return nil
}
func (r *repository) FindByID(ctx context.Context, id uint) (*User, error) {
	return nil, nil
}
func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}
func (r *repository) FindAll(ctx context.Context, user *User) ([]*User, error) {
	return nil, nil
}