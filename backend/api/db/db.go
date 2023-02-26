package db

import "os"

var Datastore *DB

// Initialize the database
func init() {
	os.MkdirAll("database/", 0700)
	Datastore = NewDB("database/restaurant.db", "user")
	Datastore.NewStore("menu")
}
