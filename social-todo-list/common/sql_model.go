package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}
