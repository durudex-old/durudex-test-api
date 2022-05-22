/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (h *Handler) isAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if status := ctx.Value("auth"); status == false || status == nil {
		return nil, &gqlerror.Error{Message: "Authorization token failed"}
	}

	return next(ctx)
}
