/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package http

import (
	"errors"
	"strings"

	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/durudex/durudex-test-api/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

const authorizationHeader string = "Authorization"

var (
	ErrAuthHeader       = errors.New("invalid auth header")
	ErrAuthTokenIsEmpty = errors.New("token is empty")
)

// Authorization HTTP middleware.
func (h *Handler) authMiddleware(ctx *fiber.Ctx) error {
	// Getting authorization header.
	header := ctx.Get(authorizationHeader)
	if header == "" {
		return ctx.Next()
	}

	// Checking header parts.
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid authorization header")
	}

	// Check the second part of the header.
	if len(headerParts[1]) == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("Authorization token is empty")
	}

	// Parsing jwt access token.
	customClaim, err := auth.Parse(headerParts[1], h.cfg.SigningKey)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Authorization token is invalid")
	}

	ctx.Context().SetUserValue(domain.UserCtx, customClaim)

	return ctx.Next()
}
