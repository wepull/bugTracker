package main

import (
	"github.com/wepull/bugTracker/database"
	"github.com/wepull/bugTracker/routes"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}
	defer db.Close()

	// Set up your router and start the server
	router := routes.SetupRouter(db)
	router.Run(":8080")
}
