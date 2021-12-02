package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Config function
func Config() {
	env, err := filepath.Abs(".env")
	if err != nil {
		log.Fatal().Err(err)
	}

	err = godotenv.Load(env)
	if err != nil {
		log.Fatal().Err(err)
	}
}
