package common

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// check if prod or local
	prod := os.Getenv("PROD")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
