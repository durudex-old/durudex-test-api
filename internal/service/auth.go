/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"

	faker "github.com/bxcodec/faker/v3"
	"github.com/durudex/durudex-test-api/internal/domain"
)

// Auth service interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error)
	RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error)
}

// Auth service structure.
type AuthService struct{}

// Creating a new auth service.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// User Sign Up.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &domain.Tokens{Access: faker.Jwt(), Refresh: faker.Password()}, nil
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	// Validate input.
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &domain.Tokens{Access: faker.Jwt(), Refresh: faker.Password()}, nil
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

	return faker.Jwt(), nil
}
