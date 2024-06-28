package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() (error, bool) {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err, true
		}
	}
	return nil, false
}
