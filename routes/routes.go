package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter initializes the main router and sets up route mappings.
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create a new gin router
	router := gin.Default()

	// Set up global middleware (if needed)

	// Set up routes
	v1 := router.Group("/api/v1")
	{
		// User endpoints
		v1.POST("/register", handlers.RegisterUser)
		v1.POST("/login", handlers.LoginUser)

		// Bug endpoints
		v1.POST("/bugs", handlers.CreateBug)
		v1.GET("/bugs/:id", handlers.GetBugByID)
		v1.GET("/bugs", handlers.GetBugs)
		v1.PUT("/bugs/:id", handlers.UpdateBug)
		v1.DELETE("/bugs/:id", handlers.DeleteBug)

		// Project endpoints
		v1.POST("/projects", handlers.CreateProject)
		v1.GET("/projects/:id", handlers.GetProjectByID)
		v1.GET("/projects", handlers.GetProjects)
		v1.PUT("/projects/:id", handlers.UpdateProject)
		v1.DELETE("/projects/:id", handlers.DeleteProject)
	}

	return router
}
