/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/domain"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQL directive for user authorization.
func (h *Handler) isAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if ctx.Value(domain.UserCtx) == nil {
		return nil, &gqlerror.Error{
			Message:    "Authorization token failed",
			Extensions: map[string]interface{}{"code": domain.CodeUnauthorized},
		}
	}

	return next(ctx)
}
