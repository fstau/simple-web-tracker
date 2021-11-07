package main

import (
	"local/tracker/db"
	"local/tracker/server"
)

func main() {
	db := db.InitDB()
	defer db.Close()
	server.InitServer(db)
}
