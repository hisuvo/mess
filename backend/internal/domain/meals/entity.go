package meals

import (
	"time"

	"messify/internal/domain/users"

	"gorm.io/gorm"
)

type MealStatus string

const (
	MealStatusActive    MealStatus = "ACTIVE"
	MealStatusPending   MealStatus = "PENDING"
	MealStatusCancelled MealStatus = "CANCELLED"
)

type Meal struct {
	gorm.Model

	TenantID  uint        `gorm:"index;not null" json:"tenant_id"`
	UserID    uint        `gorm:"index;not null" json:"user_id"`
	User      *users.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Date      time.Time   `gorm:"type:date;not null;index" json:"date"`
	Breakfast float64     `gorm:"type:decimal(5,2);default:0" json:"breakfast"`
	Lunch     float64     `gorm:"type:decimal(5,2);default:0" json:"lunch"`
	Dinner    float64     `gorm:"type:decimal(5,2);default:0" json:"dinner"`
	Total     float64     `gorm:"type:decimal(6,2);default:0" json:"total"`
	Note      string      `gorm:"type:text" json:"note,omitempty"`
	Status    MealStatus  `gorm:"type:varchar(30);not null;default:'ACTIVE'" json:"status"`
}

func (m *Meal) CalculateTotal() float64 {
	return m.Breakfast + m.Lunch + m.Dinner
}
