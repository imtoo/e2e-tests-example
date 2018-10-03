package config

import (
	"flag"
	"fmt"
	"os"

	// Automatically loads .env file
	_ "github.com/joho/godotenv/autoload"
)

// EnvDatabaseURL - URL of the server database
var EnvDatabaseURL = getEnv("DATABASE_URL")

// EnvRunSeeds - Runs seeds for DB
var EnvRunSeeds = getEnv("RUN_SEEDS") == "true"

// EnvGo - Go Environment
var EnvGo = getEnv("GO_ENV")

// EnvIsProduction - Go Environment production
var EnvIsProduction = getEnv("GO_ENV") == "production"

// EnvIsStaging - Go Environment production
var EnvIsStaging = getEnv("GO_ENV") == "staging"

// EnvIsTest - Test environment
var EnvIsTest = getIsTest()

// EnvPort - port of the server
var EnvPort = getPort()

// EnvPathAndPort - get path of the server and port
var EnvPathAndPort = getPathAndPort()

func getIsTest() bool {
	if flag.Lookup("test.v") == nil {
		return false
	}

	return true
}

func getPort() string {
	envPort := getEnv("PORT")

	if envPort != "" {
		return fmt.Sprintf(":%s", envPort)
	}

	return ":8000"
}

func getPathAndPort() string {
	if EnvIsProduction || EnvIsStaging {
		return EnvPort
	}

	// to supress Accept Incoming Network Connections warnings while launching binary in dev
	return "localhost" + EnvPort
}

func getEnv(name string) string {
	return os.Getenv(name)
}
