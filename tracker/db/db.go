package db

import (
	"database/sql"
	"fmt"
	"local/tracker/models"

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
	sql := fmt.Sprintf("INSERT INTO tracker.events (id, cts, uid, session, event, page, query, data, sts, origin, referer) VALUES (DEFAULT, '%d', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%s', '%s')",
		event.ClientTimestamp, event.User, event.Session, event.Event, event.Page, event.Query, event.Data, event.ServerTimestamp, event.Origin, event.Referer)
	_, err := DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func WriteUser(user models.User) {
	sql := fmt.Sprintf("INSERT INTO users (id, uid, ua, ua_hash, ip_addr, window_width, window_height, window_avail_width, window_avail_height, orientation, cts, sts) VALUES (DEFAULT, '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%s', '%d', '%d')", user.UserId, user.UA, user.UA_Hash, user.IPAddr, user.WindowWidth, user.WindowHeight, user.WindowAvailableWidth, user.WindowAvailableHeight, user.Orientation, user.ClientTimestamp, user.ServerTimestamp)
	_, err := DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// func WriteUA(ua models.UserAgent) (exists bool) {
// 	sql := fmt.Sprintf("INSERT INTO useragents (id, ua, ua_hash) VALUES (DEFAULT, '%s', '%s')", ua.UA, ua.UA_Hash)
// 	_, err := DB.Exec(sql)
// 	if err != nil {
// 		// Handle expected errors
// 		if err, ok := err.(*pq.Error); ok {
// 			if err.Code.Name() == "unique_violation" {
// 				return true
// 			}
// 		}
// 		panic(err)
// 	}
// 	return false
// }
