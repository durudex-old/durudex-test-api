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
	// Creating a new post.
	CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error)
	// Deleting a post.
	DeletePost(ctx context.Context, id ksuid.KSUID) (bool, error)
	// Updating a post.
	UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error)
	// Getting a post.
	Post(ctx context.Context, id ksuid.KSUID) (*domain.Post, error)
	// Getting a posts.
	Posts(ctx context.Context, sort domain.SortOptions) (*domain.PostConnection, error)
}

// Post service structure.
type PostService struct{}

// Creating a new post service.
func NewPostService() *PostService {
	return &PostService{}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return ksuid.Nil, err
	}

	return ksuid.New(), nil
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id ksuid.KSUID) (bool, error) {
	if id.IsNil() {
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
func (s *PostService) Post(ctx context.Context, id ksuid.KSUID) (*domain.Post, error) {
	if id.IsNil() {
		return nil, &gqlerror.Error{
			Message:    "Post not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return domain.NewPost(id), nil
}

// Getting a posts.
func (s *PostService) Posts(ctx context.Context, sort domain.SortOptions) (*domain.PostConnection, error) {
	var limit int

	// Check filter and last filters.
	switch {
	// Check if first and last filters is not nil.
	case sort.First != nil && sort.Last != nil:
		return nil, &gqlerror.Error{
			Message:    "Must be `first` or `last`",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	// Check if first filter is nil.
	case sort.First == nil:
		// Check if last filter is nil or set last filter.
		if sort.Last == nil {
			return nil, &gqlerror.Error{
				Message:    "Must be `first` or `last`",
				Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
			}
		} else if *sort.Last > 50 || *sort.Last < 1 {
			return nil, &gqlerror.Error{
				Message:    "`last` must not exceed 50 or be less than 1",
				Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
			}
		}

		limit = *sort.Last
	// Check if first filter is nil or set last filter.
	case *sort.First > 50 || *sort.First < 1:
		return nil, &gqlerror.Error{
			Message:    "`first` must not exceed 50 or be less than 1",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	// Set first filter.
	default:
		limit = *sort.First
	}

	if limit == 1 {
		if rand.Intn(2) == 1 {
			limit = 0
		}
	} else {
		limit = rand.Intn(limit)
	}

	nodes := make([]*domain.Post, limit)

	for i := 0; i < limit; i++ {
		nodes[i] = domain.NewPost(ksuid.New())
	}

	return &domain.PostConnection{
		Nodes:      nodes,
		TotalCount: limit + limit,
	}, nil
}
