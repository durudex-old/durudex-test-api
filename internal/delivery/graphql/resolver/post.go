package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"time"

	faker "github.com/bxcodec/faker/v3"
	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (string, error) {
	return faker.UUIDHyphenated(), nil
}

func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {
	if id == domain.FalseOther {
		return nil, &gqlerror.Error{Message: "Post not found"}
	}

	return &model.Post{
		ID:        id,
		UserID:    faker.UUIDHyphenated(),
		Text:      faker.Sentence(),
		CreatedAt: time.Unix(faker.RandomUnixTime(), 0),
		UpdatedAt: func() *time.Time {
			if rand.Intn(2) == 1 {
				updatedAt := time.Unix(faker.RandomUnixTime(), 0)
				return &updatedAt
			} else {
				return nil
			}
		}(),
	}, nil
}
