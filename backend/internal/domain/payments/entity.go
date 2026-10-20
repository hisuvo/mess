package payments

import (
	"time"

	"messify/internal/domain/users"

	"gorm.io/gorm"
)

type PaymentType string

const (
	PaymentTypeRent        PaymentType = "RENT"
	PaymentTypeUtility     PaymentType = "UTILITY"
	PaymentTypeCookSalary  PaymentType = "COOK_SALARY"
	PaymentTypeWifi        PaymentType = "WIFI"
	PaymentTypeGas         PaymentType = "GAS"
	PaymentTypeElectricity PaymentType = "ELECTRICITY"
	PaymentTypeRefund      PaymentType = "REFUND"
	PaymentTypeOther       PaymentType = "OTHER"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
)

type PaymentMethod string

const (
	PaymentMethodCash   PaymentMethod = "CASH"
	PaymentMethodBkash  PaymentMethod = "BKASH"
	PaymentMethodNagad  PaymentMethod = "NAGAD"
	PaymentMethodRocket PaymentMethod = "ROCKET"
	PaymentMethodBank   PaymentMethod = "BANK"
	PaymentMethodOther  PaymentMethod = "OTHER"
)

type Payment struct {
	gorm.Model

	TenantID      uint          `gorm:"index;not null" json:"tenant_id"`
	UserID        *uint         `gorm:"index" json:"user_id,omitempty"`
	User          *users.User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Amount        float64       `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentType   PaymentType   `gorm:"type:varchar(50);not null" json:"payment_type"`
	PaymentMethod PaymentMethod `gorm:"type:varchar(30);not null;default:'CASH'" json:"payment_method"`
	PaymentDate   time.Time     `gorm:"type:date;not null;index" json:"payment_date"`
	TransactionID string        `gorm:"type:varchar(100)" json:"transaction_id,omitempty"`
	Status        PaymentStatus `gorm:"type:varchar(30);not null;default:'COMPLETED';index" json:"status"`
	Description   string        `gorm:"type:text" json:"description,omitempty"`
}
