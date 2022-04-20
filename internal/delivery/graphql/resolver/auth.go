package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	faker "github.com/bxcodec/faker/v3"
	"github.com/durudex/durudex-test-api/internal/delivery/graphql/model"
	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (string, error) {
	if input.Username == domain.FalseOther {
		return "", &gqlerror.Error{Message: "Username already exists"}
	} else if input.Email == domain.FalseEmail {
		return "", &gqlerror.Error{Message: "Email already exists"}
	} else if input.Code == domain.FalseCode {
		return "", &gqlerror.Error{Message: "Code denied"}
	}

	return faker.UUIDHyphenated(), nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.Tokens, error) {
	if input.Username == domain.FalseOther {
		return nil, &gqlerror.Error{Message: "User not found"}
	} else if input.Password == domain.FalseOther {
		return nil, &gqlerror.Error{Message: "Password incorrect"}
	}

	return &model.Tokens{Access: faker.Jwt(), Refresh: faker.Password()}, nil
}

func (r *mutationResolver) SignOut(ctx context.Context, input model.RefreshTokenInput) (bool, error) {
	if input.Token == domain.FalseOther {
		return false, &gqlerror.Error{Message: "Session not found"}
	}

	return true, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	if input.Token == domain.FalseOther {
		return "", &gqlerror.Error{Message: "Session not found"}
	}

	return faker.Jwt(), nil
}
