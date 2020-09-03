package config

import (
	"os"
)

func getEnvOrDefault(key string, defaultValue string) string {
	result := os.Getenv(key)
	if len(result) == 0 {
		result = defaultValue
	}
	return result
}

var Key string = getEnvOrDefault("KEY", "zzxxccvvbbnnmmaa")

var PGHost string = getEnvOrDefault("PG_HOST", "127.0.0.1")
var PGPort string = getEnvOrDefault("PG_PORT", "5432")
var PGDatabase string = getEnvOrDefault("PG_DATABASE", "dashboard-dev")
var PGUser string = getEnvOrDefault("PG_USER", "dev")
var PGPassword string = getEnvOrDefault("PG_PASSWORD", "dev")
