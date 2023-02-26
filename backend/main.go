package main

import (
	"fmt"

	"goozt.org/breakstick/api"
	"goozt.org/breakstick/api/db"
)

func main() {
	server := api.NewServer("Restaurant API v1", "/api/v1")
	LoadMiddlewares(server)
	RegisterHandlers(server)
	server.Cleanup(func() error {
		fmt.Println("\rRunning cleanup tasks...")
		defer db.Datastore.Close()
		return nil
	})
	server.Listen()
}
