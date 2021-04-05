package db

import (
	"database/sql"
	"fmt"
	"local/tracker/models"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PW   = "postgres"
	DB_NAME = "tracker"
)

var DB *sql.DB

func InitDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PW, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Failed to connect to postgres")
	}
	// defer db.Close()

	DB = db
	return db
}

func WriteEvent(event models.Event) {
	sql := fmt.Sprintf("INSERT INTO events (id, event, data, cts, sts, origin, referer, ua_hash, uid) VALUES (DEFAULT, '%s', '%s', %d, '%d', '%s', '%s', '%s', '%s')", event.Event, event.Data, event.ClientTimestamp, event.ServerTimestamp, event.Origin, event.Referer, event.UserAgentHash, event.UserId)
	_, err := DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func WriteUA(ua models.UserAgent) (exists bool) {
	sql := fmt.Sprintf("INSERT INTO useragents (id, ua, ua_hash) VALUES (DEFAULT, '%s', '%s')", ua.UA, ua.UA_Hash)
	_, err := DB.Exec(sql)
	if err != nil {
		// Handle expected errors
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return true
			}
		}
		panic(err)
	}
	return false
}
