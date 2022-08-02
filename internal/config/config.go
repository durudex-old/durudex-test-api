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
	Config struct {
		HTTP    HTTPConfig    `mapstructure:"http"`
		Auth    AuthConfig    `mapstructure:"auth"`
		GraphQL GraphQLConfig `mapstructure:"graphql"`
	}

	// HTTP server config variables.
	HTTPConfig struct {
		Host string     `mapstructure:"host"`
		Port string     `mapstructure:"port"`
		Name string     `mapstructure:"name"`
		Cors CorsConfig `mapstructure:"cors"`
	}

	// CORS config variables.
	CorsConfig struct {
		Enable       bool   `mapstructure:"enable"`
		AllowOrigins string `mapstructure:"allow-origins"`
		AllowMethods string `mapstructure:"allow-methods"`
		AllowHeaders string `mapstructure:"allow-headers"`
	}

	// GraphQL config variables.
	GraphQLConfig struct {
		ComplexityLimit int `mapstructure:"complexity-limit"`
	}

	// Auth config variables.
	AuthConfig struct {
		SigningKey string
		TTL        time.Duration `mapstructure:"ttl"`
	}
)

// Creating a new config.
func NewConfig() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Parsing specified when starting the config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := viper.Unmarshal(&cfg); err != nil {
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

// Set configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment.")

	// Auth variables.
	cfg.Auth.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}
