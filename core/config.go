package core

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	BaseURL string
	Secret  string
	Database
}

type Database struct {
	URL string
}

func NewConfig() (*Config, error) {
	baseURLEnv := os.Getenv("KAAGGLE_AUTH_API_URL")
	dbURL := os.Getenv("KAAGGLE_DB_URL")
	secret := os.Getenv("KAAGGLE_SECRET")

	if baseURLEnv != "" && dbURL != "" && secret != "" {
		return &Config{
			BaseURL:  baseURLEnv,
			Database: Database{dbURL},
			Secret:   secret,
		}, nil
	}

	return nil, errors.New("Please add KAAGGLE_AUTH_API_URL, KAAGGLE_DB_URL and KAAGGLE_SECRET environmental variables")
}

func (c *Config) String() string {
	return fmt.Sprintf("Using config with the following details. URL: %s. DB_URL: %s.", c.BaseURL, c.Database.URL)
}
