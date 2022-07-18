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
		return ErrAuthHeader
	}

	// Check the second part of the header.
	if len(headerParts[1]) == 0 {
		return ErrAuthTokenIsEmpty
	}

	ctx.Context().SetUserValue("auth", true)

	return ctx.Next()
}
