/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"

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
