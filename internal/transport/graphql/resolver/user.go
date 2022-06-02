package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/durudex/durudex-test-api/internal/domain"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	return r.service.Auth.SignUp(ctx, input)
}

func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	return r.service.User.CreateVerifyEmailCode(ctx, email)
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	return r.service.User.ForgotPassword(ctx, input)
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return r.service.User.UpdateAvatar(ctx, file)
}

func (r *queryResolver) Me(ctx context.Context) (*domain.User, error) {
	return r.service.User.User(ctx, "")
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	return r.service.User.User(ctx, id)
}
