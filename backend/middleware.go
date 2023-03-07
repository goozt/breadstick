package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"goozt.org/breakstick/api"
)

func LoadMiddlewares(a *api.API) {
	a.Use(idempotency.New())
	a.Use(requestid.New())
	a.Use(recover.New())
	if config.Debug {
		a.Use(logger.New())
	}
	a.Use(pprof.New())
	a.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        config.CacheTimeout,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	a.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowOrigins: config.AllowedHosts,
	}))
}
