package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/domain"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	return r.service.Auth.SignUp(ctx, input)
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	return r.service.Auth.SignIn(ctx, input)
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error) {
	return r.service.Auth.RefreshToken(ctx, input)
}
