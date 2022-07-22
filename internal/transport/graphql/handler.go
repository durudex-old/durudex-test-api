/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import (
	"net/http"

	"github.com/durudex/durudex-test-api/internal/service"
	"github.com/durudex/durudex-test-api/internal/transport/graphql/generated"
	"github.com/durudex/durudex-test-api/internal/transport/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
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
	// GraphQL config.
	config := generated.Config{
		Resolvers:  resolver.NewResolver(h.service),
		Directives: generated.DirectiveRoot{IsAuth: h.isAuth},
	}

	// User posts complexity.
	config.Complexity.User.Posts = func(childComplexity int, first, last *int) int {
		switch {
		case first != nil:
			return childComplexity * *first
		case last != nil:
			return childComplexity * *last
		default:
			return 0
		}
	}

	// Post author complexity.
	config.Complexity.Post.Author = func(childComplexity int) int { return childComplexity * 2 }

	// Creating a new graphql handler.
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	// Set graphql complexity limit.
	handler.Use(extension.FixedComplexityLimit(500))

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// GraphQL playground handler.
func (h *Handler) PlaygroundHandler() http.HandlerFunc {
	return playground.Handler("GraphQL", "/query")
}
