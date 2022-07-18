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

	"github.com/segmentio/ksuid"
)

// Auth service interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error)
	RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error)
	CreateSession() (*domain.Tokens, error)
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

	return s.CreateSession()
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return s.CreateSession()
}

// Creating a new fake user session.
func (s *AuthService) CreateSession() (*domain.Tokens, error) {
	// Generating a new jwt access token
	accessToken, err := auth.GenerateAccessToken(ksuid.New().String(), s.cfg.SigningKey, s.cfg.TTL)
	if err != nil {
		return nil, err
	}

	// Generating a new refresh token.
	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{Access: accessToken, Refresh: refreshToken}, nil
}

// User Sign Out.
func (s *AuthService) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return false, err
	}

	return true, nil
}

// Refresh user access token token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return "", err
	}

	return auth.GenerateAccessToken(ksuid.New().String(), s.cfg.SigningKey, s.cfg.TTL)
}
