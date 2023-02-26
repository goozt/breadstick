package main

import (
	"goozt.org/breakstick/api"
	"goozt.org/breakstick/api/menu"
)

func RegisterHandlers(a *api.API) {
	menu.HandleRoutes(a.Router)
}
