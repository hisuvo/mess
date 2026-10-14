package bazars

import (
	"time"

	"messify/internal/domain/users"

	"gorm.io/gorm"
)

type BazarStatus string

const (
	BazarStatusPending  BazarStatus = "PENDING"
	BazarStatusApproved BazarStatus = "APPROVED"
	BazarStatusRejected BazarStatus = "REJECTED"
)

type Bazar struct {
	gorm.Model

	TenantID     uint        `gorm:"index;not null" json:"tenant_id"`
	UserID       uint        `gorm:"index;not null" json:"user_id"`
	User         *users.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ItemName     string      `gorm:"type:varchar(255)" json:"item_name"`
	Quantity     string      `gorm:"type:varchar(100)" json:"quantity,omitempty"`
	TotalAmount  float64     `gorm:"type:decimal(10,2);default:0" json:"total_amount"`
	SpentAmount  float64     `gorm:"type:decimal(10,2);default:0" json:"spent_amount"`
	PurchaseDate time.Time   `gorm:"type:date;not null;index" json:"purchase_date"`
	Description  string      `gorm:"type:text" json:"description,omitempty"`
	Status       BazarStatus `gorm:"type:varchar(30);not null;default:'APPROVED';index" json:"status"`
}

type BazarList struct {
	Bazar []Bazar `json:"bazar"`
}