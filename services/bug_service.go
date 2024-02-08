package services

import (
	"github.com/wepull/bugTracker/models"
	"github.com/wepull/bugTracker/repositories"
)

// BugService handles bug-related operations.
type BugService struct {
	bugRepo repositories.BugRepository
}

// NewBugService creates a new instance of the BugService.
func NewBugService(bugRepo repositories.BugRepository) *BugService {
	return &BugService{bugRepo: bugRepo}
}

// CreateBug creates a new bug.
func (s *BugService) CreateBug(bug *models.Bug) error {
	// Perform bug data validation.
	// Set the initial status of the bug.
	// Call the bug repository to store the bug in the database.

	return s.bugRepo.CreateBug(bug)
}

// GetBugByID gets bug details by ID.
func (s *BugService) GetBugByID(id uint) (*models.Bug, error) {
	// Check if the bug with the given ID exists in the database.
	// Call the bug repository to fetch the bug details.

	return s.bugRepo.GetBugByID(id)
}
