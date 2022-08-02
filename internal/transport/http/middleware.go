/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package http

import (
	"strings"

	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/durudex/durudex-test-api/pkg/auth"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofiber/fiber/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const authorizationHeader string = "Authorization"

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
		return ctx.Status(fiber.StatusBadRequest).JSON(
			graphql.Response{
				Errors: []*gqlerror.Error{
					{
						Message:    "Invalid authorization header",
						Extensions: map[string]interface{}{"code": domain.CodeUnauthorized},
					},
				},
			},
		)
	}

	// Parsing jwt access token.
	customClaim, err := auth.Parse(headerParts[1], h.signingKey)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			graphql.Response{
				Errors: []*gqlerror.Error{
					{
						Message:    "Authorization token is invalid",
						Extensions: map[string]interface{}{"code": domain.CodeUnauthorized},
					},
				},
			},
		)
	}

	ctx.Context().SetUserValue(domain.UserCtx, customClaim)

	return ctx.Next()
}
