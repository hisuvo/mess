package deposits

import (
	"time"

	"messify/internal/domain/users"

	"gorm.io/gorm"
)

type DepositStatus string

const (
	DepositStatusPending  DepositStatus = "PENDING"
	DepositStatusApproved DepositStatus = "APPROVED"
	DepositStatusRejected DepositStatus = "REJECTED"
)

type DepositMethod string

const (
	DepositMethodCash   DepositMethod = "CASH"
	DepositMethodBkash  DepositMethod = "BKASH"
	DepositMethodNagad  DepositMethod = "NAGAD"
	DepositMethodRocket DepositMethod = "ROCKET"
	DepositMethodBank   DepositMethod = "BANK"
	DepositMethodOther  DepositMethod = "OTHER"
)

type Deposit struct {
	gorm.Model

	TenantID      uint          `gorm:"index;not null" json:"tenant_id"`
	UserID        uint          `gorm:"index;not null" json:"user_id"`
	User          *users.User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Amount        float64       `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentDate   time.Time     `gorm:"type:date;not null;index" json:"payment_date"`
	Method        DepositMethod `gorm:"type:varchar(30);not null;default:'CASH'" json:"method"`
	TransactionID string        `gorm:"type:varchar(100)" json:"transaction_id,omitempty"`
	Status        DepositStatus `gorm:"type:varchar(30);not null;default:'APPROVED';index" json:"status"`
	Note          string        `gorm:"type:text" json:"note,omitempty"`
}

