package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
)

func (r *mutationResolver) GetCodeByEmail(ctx context.Context, input model.GetCodeByEmailInput) (*model.Status, error) {
	return &model.Status{Status: true}, nil
}
