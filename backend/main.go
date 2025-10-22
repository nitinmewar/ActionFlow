package main

import (
	"log"
	"orbit/cmd/database"
	"orbit/cmd/env"
	"orbit/cmd/server"
	"orbit/cmd/utils"
)

func main() {
	// load configs
	env.Load()

	// connect DB
	db, _ := database.Connection()

	// migrate db
	utils.Migrate()

	// setup server
	port := env.Port.GetValue()
	if port == "" {
		port = "8000"
	}

	r := server.SetupRouter(db)
	log.Fatal(r.Run(":" + port))
}
