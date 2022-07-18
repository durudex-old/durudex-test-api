package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/domain"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	return r.service.Post.CreatePost(ctx, input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	return r.service.Post.DeletePost(ctx, id)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	return r.service.Post.UpdatePost(ctx, input)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*domain.Post, error) {
	return r.service.Post.Post(ctx, id)
}
