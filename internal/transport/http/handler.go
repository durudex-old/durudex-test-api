/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package http

import (
	"github.com/durudex/durudex-test-api/internal/service"
	"github.com/durudex/durudex-test-api/internal/transport/graphql"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// HTTP handler structure.
type Handler struct{ service *service.Service }

// Creating a new http handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	// Creating a new graphql handler.
	graphql := graphql.NewHandler(h.service)

	router.Use(h.authMiddleware)

	router.Get("/", adaptor.HTTPHandlerFunc(graphql.PlaygroundHandler()))
	router.Post("/query", adaptor.HTTPHandlerFunc(graphql.GraphqlHandler()))
}
