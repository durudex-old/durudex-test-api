package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/durudex/durudex-test-api/internal/domain"
)

// CreateVerifyEmailCode is the resolver for the createVerifyEmailCode field.
func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	return r.service.User.CreateVerifyEmailCode(ctx, email)
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	return r.service.User.ForgotPassword(ctx, input)
}

// UpdateAvatar is the resolver for the updateAvatar field.
func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return r.service.User.UpdateAvatar(ctx, file)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*domain.User, error) {
	return r.service.User.User(ctx, "")
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	return r.service.User.User(ctx, id)
}
