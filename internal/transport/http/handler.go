/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package http

import (
	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/transport/graphql"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// HTTP handler structure.
type Handler struct {
	config     *config.HTTPConfig
	graphql    *graphql.Handler
	signingKey string
}

// Creating a new HTTP handler.
func NewHandler(config *config.HTTPConfig, graphql *graphql.Handler, signingKey string) *Handler {
	return &Handler{config: config, graphql: graphql, signingKey: signingKey}
}

// Initialize http middleware.
func (h *Handler) InitMiddleware(router fiber.Router) {
	if h.config.Cors.Enable {
		// CORS configuration.
		corsConfig := cors.Config{
			AllowOrigins: h.config.Cors.AllowOrigins,
			AllowMethods: h.config.Cors.AllowMethods,
			AllowHeaders: h.config.Cors.AllowHeaders,
		}

		// Set CORS middleware.
		router.Use(cors.New(corsConfig))
	}

	// Set default headers.
	router.Use(h.authMiddleware)
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	// Ping pong route.
	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	// GraphQL routes.
	router.Get("/", adaptor.HTTPHandlerFunc(h.graphql.PlaygroundHandler()))
	router.Post("/query", adaptor.HTTPHandlerFunc(h.graphql.GraphqlHandler()))
}
