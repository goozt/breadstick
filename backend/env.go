package main

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	godotenv.Load()
	if env, exists := os.LookupEnv("ENVIRONMENT"); exists {
		config.Debug = strings.ToLower(env) != "production"
	}

	if config.Debug {
		if debug, exists := os.LookupEnv("DEBUG"); exists {
			config.Debug = strings.ToLower(debug) != "false"
		}
	}

	if prefix, exists := os.LookupEnv("API_PREFIX"); exists {
		if strings.HasPrefix(prefix, "/") {
			config.APIPrefix = prefix
		}
	}

}
