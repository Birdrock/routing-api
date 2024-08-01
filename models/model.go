package models

import "time"

type Model struct {
	Guid      string    `gorm:"primary_key; size:36" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
