package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wepull/bugTracker/models"
)

// ProjectRepository provides methods to interact with project data in the database.
type ProjectRepository struct {
	db *gorm.DB
}

// NewProjectRepository creates a new instance of ProjectRepository.
func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// CreateProject creates a new project in the database.
func (r *ProjectRepository) CreateProject(project *models.Project) error {
	return r.db.Create(project).Error
}

// GetProjectByID gets a project by ID from the database.
func (r *ProjectRepository) GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := r.db.Where("id = ?", id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}
