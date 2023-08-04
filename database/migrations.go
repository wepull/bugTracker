package database

import (
	"github.com/jinzhu/gorm"
	"github.com/wepull/bugTracker/models"
)

// MigrateDB performs the database migration.
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Bug{}, &models.Project{}, &models.User{})
}
