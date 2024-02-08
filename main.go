package main

import (
	"fmt"

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
	err = routes.StartHTTPSvr()
	if err != nil {
		fmt.Printf("Error Starting the Bug Tracker Backend Server: %v\n", err)
		return
	}
}
