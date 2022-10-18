/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import (
	"net/http"

	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/service"
	"github.com/durudex/durudex-test-api/internal/transport/graphql/generated"
	"github.com/durudex/durudex-test-api/internal/transport/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
)

// GraphQL handler structure.
type Handler struct {
	service *service.Service
	config  *config.GraphQLConfig
}

// Creating a new graphql handler.
func NewHandler(service *service.Service, config *config.GraphQLConfig) *Handler {
	return &Handler{service: service, config: config}
}

// GraphQL handler.
func (h *Handler) GraphqlHandler() http.HandlerFunc {
	// GraphQL config.
	config := generated.Config{
		Resolvers:  resolver.NewResolver(h.service),
		Directives: generated.DirectiveRoot{IsAuth: h.isAuth},
	}

	// Setting the complexity of the query.
	setComplexity(&config.Complexity)

	// Creating a new graphql handler.
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	// Set graphql complexity limit.
	handler.Use(extension.FixedComplexityLimit(h.config.ComplexityLimit))

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// GraphQL playground handler.
func (h *Handler) PlaygroundHandler() http.HandlerFunc {
	return playground.Handler("Durudex Test API", "/query")
}
