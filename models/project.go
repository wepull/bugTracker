package models

import (
	"time"
)

// Project represents the data model for a project.
type Project struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `gorm:"type:date" json:"start_date"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}
