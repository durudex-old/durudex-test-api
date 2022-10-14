/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"
	"math/rand"

	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/durudex/durudex-test-api/pkg/auth"
	"github.com/durudex/go-refresh"

	"github.com/segmentio/ksuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Auth service interface.
type Auth interface {
	// User Sign Up.
	SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error)
	// User Sign In.
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	// User Sign Out.
	SignOut(ctx context.Context, input domain.SessionCredInput) (bool, error)
	// Refresh user access token token.
	RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error)
	// Creating a new fake user session.
	CreateSession(userId ksuid.KSUID) (*domain.Tokens, error)
	// Deleting a user session.
	DeleteSession(ctx context.Context, input domain.DeleteSessionInput) (bool, error)
	// Getting a user session.
	Session(ctx context.Context, id ksuid.KSUID) (*domain.Session, error)
	// Getting a user sessions.
	Sessions(ctx context.Context, sort domain.SortOptions) (*domain.SessionConnection, error)
}

// Auth service structure.
type AuthService struct{ cfg *config.AuthConfig }

// Creating a new auth service.
func NewAuthService(cfg *config.AuthConfig) *AuthService {
	return &AuthService{cfg: cfg}
}

// User Sign Up.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return s.CreateSession(ksuid.New())
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return s.CreateSession(ksuid.New())
}

// Creating a new fake user session.
func (s *AuthService) CreateSession(userId ksuid.KSUID) (*domain.Tokens, error) {
	// Generating a new jwt access token
	accessToken, err := auth.GenerateAccessToken(userId.String(), s.cfg.SigningKey, s.cfg.TTL)
	if err != nil {
		return nil, err
	}

	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		return nil, err
	}
	refreshToken := r.Token(ksuid.New().String(), userId.String())

	return &domain.Tokens{Access: accessToken, Refresh: refreshToken}, nil
}

// User Sign Out.
func (s *AuthService) SignOut(ctx context.Context, input domain.SessionCredInput) (bool, error) {
	// Parsing refresh token string.
	if _, err := refresh.Parse(input.Refresh); err != nil {
		return false, err
	}

	return true, nil
}

// Refresh user access token token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error) {
	// Parsing refresh token string.
	if _, err := refresh.Parse(input.Refresh); err != nil {
		return "", err
	}

	return auth.GenerateAccessToken(ksuid.New().String(), s.cfg.SigningKey, s.cfg.TTL)
}

// Deleting a user session.
func (s *AuthService) DeleteSession(ctx context.Context, input domain.DeleteSessionInput) (bool, error) {
	if _, err := refresh.Parse(input.Refresh); err != nil {
		return false, err
	} else if input.Id.IsNil() {
		return false, &gqlerror.Error{
			Message:    "User not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return true, nil
}

// Getting a user session.
func (s *AuthService) Session(ctx context.Context, id ksuid.KSUID) (*domain.Session, error) {
	if id.IsNil() {
		return nil, &gqlerror.Error{
			Message:    "Session not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return domain.NewSession(id), nil
}

// Getting a user sessions.
func (s *AuthService) Sessions(ctx context.Context, sort domain.SortOptions) (*domain.SessionConnection, error) {
	// Validating query sort options.
	limit, err := sort.Validate()
	if err != nil {
		return nil, err
	}

	if limit == 1 {
		if rand.Intn(2) == 1 {
			limit = 0
		}
	} else {
		limit = rand.Intn(limit)
	}

	nodes := make([]*domain.Session, limit)

	for i := 0; i < limit; i++ {
		nodes[i] = domain.NewSession(ksuid.New())
	}

	return &domain.SessionConnection{
		Nodes:      nodes,
		TotalCount: limit + limit,
	}, nil
}
