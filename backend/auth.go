package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/joho/godotenv"
)

var DEFAULT_ENTROPY = 32

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	if config.APIKey != "" {
		hashedAPIKey := sha256.Sum256([]byte(config.APIKey))
		hashedKey := sha256.Sum256([]byte(key))

		if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
			return true, nil
		}
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func GenerateSecureToken(length ...int) string {

	byteLength := DEFAULT_ENTROPY
	if len(length) > 0 && length[0] != 0 {
		byteLength = length[0]
	}

	println(byteLength)
	token_byte := make([]byte, byteLength)
	if _, err := rand.Read(token_byte); err != nil {
		return err.Error()
	}

	encoded := base64.StdEncoding.EncodeToString(token_byte)
	encoded = strings.ReplaceAll(encoded, "+", "")
	encoded = strings.ReplaceAll(encoded, "/", "")
	return strings.ReplaceAll(encoded, "=", "")
}

func LoadKey() string {
	godotenv.Load()
	if apiToken, exists := os.LookupEnv("API_TOKEN"); exists {
		return apiToken
	}
	log.Fatal("API token not found")
	return ""
}

func TokenCmd() {
	argLen := len(os.Args)
	if argLen > 1 && os.Args[1] == "-t" {
		length := 0
		if argLen > 2 && os.Args[2] != "" {
			length, _ = strconv.Atoi(os.Args[2])
		}

		fmt.Println("VITE_API_TOKEN=" + GenerateSecureToken(length))
		os.Exit(0)
	}
}
