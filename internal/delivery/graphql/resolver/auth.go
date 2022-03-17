package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"

	faker "github.com/bxcodec/faker/v3"
	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
)

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (uint64, error) {
	return rand.Uint64(), nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.Tokens, error) {
	return &model.Tokens{Access: faker.Jwt(), Refresh: faker.Password()}, nil
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input model.RefreshTokensInput) (*model.Tokens, error) {
	return &model.Tokens{Access: faker.Jwt(), Refresh: faker.Password()}, nil
}
