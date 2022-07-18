/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default config path.
const defaultConfigPath string = "configs/main"

type (
	// Config variables.
	Config struct{ Auth AuthConfig }

	// Auth config variables.
	AuthConfig struct {
		SigningKey string
		TTL        time.Duration `mapstructure:"ttl"`
	}
)

// Initialize config.
func Init() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Parsing specified when starting the config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Set configurations from environment.
	setFromEnv(&cfg)

	return &cfg, nil
}

// Parsing specified when starting the config file.
func parseConfigFile() error {
	// Get config path variable.
	configPath := os.Getenv("CONFIG_PATH")

	// Check is config path variable empty.
	if configPath == "" {
		configPath = defaultConfigPath
	}

	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	dir, file := filepath.Split(configPath)

	viper.AddConfigPath(dir)
	viper.SetConfigName(file)

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal auth keys.
	return viper.UnmarshalKey("auth", &cfg.Auth)
}

// Set configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment.")

	// Auth variables.
	cfg.Auth.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}
