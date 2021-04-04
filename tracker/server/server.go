package server

import "database/sql"

func InitServer(db *sql.DB) {
	r := NewRouter(db)
	r.Run()
}
