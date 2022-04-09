package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"time"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
	"github.com/durudex/durudex-test-api/internal/domain"

	"github.com/bxcodec/faker/v3"
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

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	if id == domain.FalseOther {
		return nil, &gqlerror.Error{Message: "User not found"}
	}

	return &model.User{
		ID:        id,
		Username:  faker.Username(),
		JoinedIn:  time.Unix(faker.RandomUnixTime(), 0),
		LastVisit: time.Unix(faker.RandomUnixTime(), 0),
		Verified:  rand.Intn(2) == 1,
		AvatarURL: nil,
	}, nil
}
