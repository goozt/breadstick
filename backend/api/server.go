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
	App         *fiber.App
	cleanupFunc func() error
	Router      fiber.Router
}

type Config struct {
	AppName      string
	APIPrefix    string
	AllowedHosts string
	Debug        bool
	CacheTimeout time.Duration
}

// Create new API server
func NewServer(config Config, fiberConfig ...fiber.Config) *API {

	var newConfig fiber.Config
	newConfig.ReadTimeout = 10 * time.Second

	if len(fiberConfig) > 0 {
		newConfig = fiberConfig[0]
	}

	newConfig.AppName = config.AppName
	app := fiber.New(newConfig)
	prefixedRouter := app.Group(config.APIPrefix)

	return &API{
		App: app,
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
		api.App.Hooks().OnShutdown(api.cleanupFunc)
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
		return api.App.Listen(":" + port)
	})

	// Shutdown gracefully
	g.Go(func() error {
		<-gCtx.Done()
		return api.App.ShutdownWithTimeout(api.App.Config().ReadTimeout)
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Printf("\rerror: %s \n", err)
	}

}

// Add middlewares
func (api *API) Use(middleware func(*fiber.Ctx) error) {
	api.App.Use(middleware)
}
