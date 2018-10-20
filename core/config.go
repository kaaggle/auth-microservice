package core

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	BaseURL        string
	Secret         string
	CasbinConfPath string
	Database
}

type Database struct {
	URL string
}

func NewConfig() (*Config, error) {
	baseURLEnv := "localhost:3300"
	dbURL := "mongodb://church-adoration:church-adoration1@ds125693.mlab.com:25693/church-adoration"
	secret := "MySECRET"
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	casbinConfPath := wd + "/authorization/conf/"

	if baseURLEnv != "" && dbURL != "" && secret != "" {
		return &Config{
			BaseURL:        baseURLEnv,
			Database:       Database{dbURL},
			CasbinConfPath: casbinConfPath,
			Secret:         secret,
		}, nil
	}

	return nil, errors.New("Please add config variables")
}

func (c *Config) String() string {
	return fmt.Sprintf("Using config with the following details. URL: %s. DB_URL: %s.", c.BaseURL, c.Database.URL)
}
