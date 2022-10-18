/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/config"
	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/durudex/durudex-test-api/pkg/auth"
	"github.com/durudex/go-refresh"

	"github.com/segmentio/ksuid"
)

// Auth service interface.
type Auth interface {
	// User Sign Up.
	SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error)
	// User Sign In.
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	// Refresh user access token token.
	RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error)
}

// Auth service structure.
type AuthService struct {
	session Session
	cfg     *config.AuthConfig
}

// Creating a new auth service.
func NewAuthService(session Session, cfg *config.AuthConfig) *AuthService {
	return &AuthService{session: session, cfg: cfg}
}

// User Sign Up.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return s.session.Create(ksuid.New())
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return s.session.Create(ksuid.New())
}

// Refresh user access token token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error) {
	// Parsing refresh token string.
	if _, err := refresh.Parse(input.Refresh); err != nil {
		return "", err
	}

	return auth.GenerateAccessToken(ksuid.New().String(), s.cfg.SigningKey, s.cfg.TTL)
}
