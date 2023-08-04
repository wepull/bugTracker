package services

import (
	"github.com/wepull/bugTracker/models"
	"github.com/wepull/bugTracker/repositories"
)

// ProjectService handles project-related operations.
type ProjectService struct {
	projectRepo repositories.ProjectRepository
}

// NewProjectService creates a new instance of the ProjectService.
func NewProjectService(projectRepo repositories.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

// CreateProject creates a new project.
func (s *ProjectService) CreateProject(project *models.Project) error {
	// Perform project data validation.
	// Call the project repository to store the project in the database.

	return s.projectRepo.CreateProject(project)
}

// GetProjectByID gets project details by ID.
func (s *ProjectService) GetProjectByID(id uint) (*models.Project, error) {
	// Check if the project with the given ID exists in the database.
	// Call the project repository to fetch the project details.

	return s.projectRepo.GetProjectByID(id)
}

// Implement other project-related service functions as needed.
