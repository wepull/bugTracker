package models

import (
	"time"
)

// Bug represents the data model for a bug.
type Bug struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Status      string    `gorm:"not null" json:"status"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}
