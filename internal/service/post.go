/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"
	"math/rand"

	"github.com/durudex/durudex-test-api/internal/domain"

	"github.com/segmentio/ksuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Post service interface.
type Post interface {
	CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error)
	DeletePost(ctx context.Context, id string) (bool, error)
	UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error)
	Post(ctx context.Context, id string) (*domain.Post, error)
	Posts(ctx context.Context, first, last *int) (*domain.PostConnection, error)
}

// Post service structure.
type PostService struct{}

// Creating a new post service.
func NewPostService() *PostService {
	return &PostService{}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return "", err
	}

	return ksuid.New().String(), nil
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id string) (bool, error) {
	if id == "0" {
		return false, &gqlerror.Error{
			Message:    "Post not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return true, nil
}

// Updating a post.
func (s *PostService) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return false, err
	}

	return true, nil
}

// Getting a post.
func (s *PostService) Post(ctx context.Context, id string) (*domain.Post, error) {
	if id == "0" {
		return nil, &gqlerror.Error{
			Message:    "Post not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return domain.NewPost(id), nil
}

// Getting a post connection.
func (s *PostService) Posts(ctx context.Context, first, last *int) (*domain.PostConnection, error) {
	var filter int

	// Check filter and last filters.
	switch {
	// Check if first and last filters is not nil.
	case first != nil && last != nil:
		return nil, &gqlerror.Error{
			Message:    "Must be `first` or `last`",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	// Check if first filter is nil.
	case first == nil:
		// Check if last filter is nil or set last filter.
		if last == nil {
			return nil, &gqlerror.Error{
				Message:    "Must be `first` or `last`",
				Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
			}
		} else if *last > 50 || *last < 1 {
			return nil, &gqlerror.Error{
				Message:    "`last` must not exceed 50 or be less than 1",
				Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
			}
		}

		filter = *last
	// Check if first filter is nil or set last filter.
	case *first > 50 || *first < 1:
		return nil, &gqlerror.Error{
			Message:    "`first` must not exceed 50 or be less than 1",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	// Set first filter.
	default:
		filter = *first
	}

	n := rand.Intn(filter)
	posts := make([]*domain.Post, n)

	for i := 0; i < n; i++ {
		posts[i] = domain.NewPost(ksuid.New().String())
	}

	return &domain.PostConnection{Nodes: posts}, nil
}
