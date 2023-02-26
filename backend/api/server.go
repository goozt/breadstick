package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

type API struct {
	app         *fiber.App
	cleanupFunc func() error
	Router      fiber.Router
}

// Create new API server
func NewServer(name, prefix string, config ...fiber.Config) *API {

	var newConfig fiber.Config
	newConfig.ReadTimeout = 10 * time.Second

	if len(config) > 0 {
		newConfig = config[0]
	}

	newConfig.AppName = name
	app := fiber.New(newConfig)
	prefixedRouter := app.Group(prefix)

	return &API{
		app: app,
		cleanupFunc: func() error {
			fmt.Println("\rServer stopped")
			return nil
		},
		Router: prefixedRouter,
	}
}

// Set Cleanup function for server
func (api *API) Cleanup(fn func() error) {
	api.cleanupFunc = fn
}

// Run listener
func (api *API) Listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Hook Cleanup function
	if api.cleanupFunc != nil {
		api.app.Hooks().OnShutdown(api.cleanupFunc)
	}

	// Listen to interrupts with context
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	// Run listener
	g.Go(func() error {
		return api.app.Listen(":" + port)
	})

	// Shutdown gracefully
	g.Go(func() error {
		<-gCtx.Done()
		return api.app.ShutdownWithTimeout(api.app.Config().ReadTimeout)
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Printf("\rerror: %s \n", err)
	}

}

// Add middlewares
func (api *API) Use(middleware func(*fiber.Ctx) error) {
	api.app.Use(middleware)
}
