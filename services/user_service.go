package services

import (
	"github.com/wepull/bugTracker/models"
	"github.com/wepull/bugTracker/repositories"
)

// UserService handles user-related operations.
type UserService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new instance of the UserService.
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user *models.User) error {
	// Perform user data validation.
	// Check for duplicate usernames, email, etc.
	// Hash the user's password before storing it in the database.

	return s.userRepo.CreateUser(user)
}

// AuthenticateUser authenticates a user.
func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	// Validate the user's credentials.
	// Check if the provided username exists and if the password is correct.

	return s.userRepo.GetUserByUsername(username)
}

// Implement other user-related service functions as needed.
