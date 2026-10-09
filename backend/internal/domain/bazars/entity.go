package bazars

import (
	"time"
)

type Bazar struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `json:"name"`
	TotalAmount float64   `json:"total_amount"`
	SpentAmount float64   `json:"spent_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Mile        *miles.Mile `gorm:"foreignKey:MileID"`
	MileID      int       `json:"mile_id"`
}

type BazarList struct {
	Bazar []Bazar `json:"bazar"`
}