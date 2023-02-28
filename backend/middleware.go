package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/keyauth/v2"
	"goozt.org/breakstick/api"

	"github.com/joho/godotenv"
)

var (
	apiKey string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	if apiToken, exists := os.LookupEnv("VITE_API_TOKEN"); exists {
		apiKey = apiToken
	}
}

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	if apiKey != "" {
		hashedAPIKey := sha256.Sum256([]byte(apiKey))
		hashedKey := sha256.Sum256([]byte(key))

		if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
			return true, nil
		}
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func LoadMiddlewares(a *api.API) {
	a.Use(idempotency.New())
	a.Use(requestid.New())
	a.Use(recover.New())
	a.Use(logger.New())
	a.Use(pprof.New())
	a.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	a.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
		AllowOrigins:     "https://breadstick.goozt.org, http://localhost:5173",
	}))
	a.Use(keyauth.New(keyauth.Config{
		Validator: validateAPIKey,
	}))
}
