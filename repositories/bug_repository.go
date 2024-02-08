package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wepull/bugTracker/models"
)

// BugRepository provides methods to interact with bug data in the database.
type BugRepository struct {
	db *gorm.DB
}

// NewBugRepository creates a new instance of BugRepository.
func NewBugRepository(db *gorm.DB) *BugRepository {
	return &BugRepository{db: db}
}

// CreateBug creates a new bug in the database.
func (r *BugRepository) CreateBug(bug *models.Bug) error {
	return r.db.Create(bug).Error
}

// GetBugByID gets a bug by ID from the database.
func (r *BugRepository) GetBugByID(id uint) (*models.Bug, error) {
	var bug models.Bug
	if err := r.db.Where("id = ?", id).First(&bug).Error; err != nil {
		return nil, err
	}
	return &bug, nil
}
