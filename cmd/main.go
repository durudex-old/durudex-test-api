/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/service"
	"github.com/durudex/durudex-test-api/internal/transport/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Initialize application.
func init() {
	// Set logger mode.
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// A function that running the application.
func main() {
	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Error().Err(err).Msg("error initialize config")
	}

	// Creating a new service.
	service := service.NewService(cfg)
	// Creating a new http handler.
	handler := http.NewHandler(service, &cfg.Auth)

	// Create a new server.
	srv := http.NewServer(&cfg.HTTP, handler)

	// Run server.
	go srv.Run()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stopping server.
	srv.Stop()

	log.Info().Msg("Durudex Test API stopping!")
}
