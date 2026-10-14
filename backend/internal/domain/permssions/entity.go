package permssions

import "gorm.io/gorm"

type Permission struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Slug        string `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	Module      string `gorm:"type:varchar(50);index" json:"module"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}
