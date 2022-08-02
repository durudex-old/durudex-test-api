/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package http

import (
	"github.com/durudex/durudex-test-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// HTTP server structure.
type Server struct {
	app     *fiber.App
	handler *Handler
	config  *config.HTTPConfig
}

// Creating a new application http server.
func NewServer(config *config.HTTPConfig, handler *Handler) *Server {
	return &Server{
		app:     fiber.New(fiber.Config{AppName: config.Name}),
		handler: handler,
		config:  config,
	}
}

// Running application http server.
func (s *Server) Run() {
	log.Debug().Msg("Running http server...")

	// Initialize http middleware.
	s.handler.InitMiddleware(s.app)

	// Initialize http routes.
	s.handler.InitRoutes(s.app)

	addr := s.config.Host + ":" + s.config.Port

	// Listen serves HTTP requests from the given addr.
	if err := s.app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("failed to start http server")
	}
}

// Stopping application http server.
func (s *Server) Stop() {
	log.Info().Msg("Stopping http server...")

	if err := s.app.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop http server")
	}
}
