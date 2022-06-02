package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/domain"
)

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	return r.service.Auth.SignIn(ctx, input)
}

func (r *mutationResolver) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	return r.service.Auth.SignOut(ctx, input)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	return r.service.Auth.RefreshToken(ctx, input)
}
