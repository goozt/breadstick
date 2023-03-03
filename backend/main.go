package main

import (
	"fmt"
	"time"

	"goozt.org/breakstick/api"
	"goozt.org/breakstick/api/db"
)

var config = api.Config{
	AppName:      "Restaurant API v1",
	APIPrefix:    "/api/v1",
	AllowedHosts: "https://breadstick.goozt.org, http://localhost:5173, http://localhost:4173",
	Debug:        false,
	CacheTimeout: 30 * time.Second,
}

func main() {
	TokenCmd()
	config.APIKey = LoadKey()
	server := api.NewServer(config)
	LoadMiddlewares(server)
	RegisterHandlers(server)
	server.Cleanup(func() error {
		fmt.Println("\rRunning cleanup tasks...")
		defer db.Datastore.Close()
		return nil
	})
	server.Listen()
}
