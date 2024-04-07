package main

import (
	"knockknock/api"
	"knockknock/db"
)

func main() {

	// 1. Connect with DB
	db_conn := db.InitDB()

	// 2. Pass DB to API server
	server := api.NewAPIServer(":8080", db_conn)

	// 3. Run server
	server.Serve()

}
