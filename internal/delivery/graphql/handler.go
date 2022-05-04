/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import (
	"net/http"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql/generated"
	"github.com/durudex/durudex-test-api/internal/delivery/graphql/resolver"
	"github.com/durudex/durudex-test-api/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// GraphQL handler structure.
type Handler struct{ service *service.Service }

// Creating a new graphql handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// GraphQL handler.
func (h *Handler) GraphqlHandler() http.HandlerFunc {
	config := generated.Config{
		Resolvers:  resolver.NewResolver(h.service),
		Directives: generated.DirectiveRoot{IsAuth: h.isAuth},
	}

	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// GraphQL playground handler.
func (h *Handler) PlaygroundHandler() http.HandlerFunc {
	return playground.Handler("GraphQL", "/query")
}
