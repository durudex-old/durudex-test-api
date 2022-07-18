/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"os"

	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/service"
	"github.com/durudex/durudex-test-api/internal/transport/http"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init config")
	}

	// Crating a new fiber app.
	app := fiber.New(fiber.Config{})

	// Use cors middleware.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// Creating a new service.
	service := service.NewService(cfg)
	// Creating a new http handler.
	handler := http.NewHandler(service, &cfg.Auth)

	// Initialize http routes.
	handler.InitRoutes(app)

	port := os.Getenv("API_PORT")

	log.Printf("Server is run it ':%s'", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running http server")
	}
}
