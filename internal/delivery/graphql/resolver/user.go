package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) ForgotPassword(ctx context.Context, input model.ForgotPasswordInput) (bool, error) {
	if input.Email == domain.FalseEmail {
		return false, &gqlerror.Error{Message: "Email denied"}
	} else if input.Code == domain.FalseCode {
		return false, &gqlerror.Error{Message: "Code denied"}
	}

	return true, nil
}
