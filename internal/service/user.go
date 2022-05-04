/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import (
	"context"

	"github.com/durudex/durudex-test-api/internal/domain"

	"github.com/99designs/gqlgen/graphql"
)

// User service interface.
type User interface {
	CreateVerifyEmailCode(ctx context.Context, email string) (bool, error)
	ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error)
	UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error)
	User(ctx context.Context, id string) (*domain.User, error)
}

// User service structure.
type UserService struct{}

// Creating a new user service.
func NewUserService() *UserService {
	return &UserService{}
}

// Creating a new user verification email code.
func (s *UserService) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	return true, nil
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	return true, nil
}

// Update user avatar.
func (s *UserService) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return "", nil
}

// Getting a user.
func (s *UserService) User(ctx context.Context, id string) (*domain.User, error) {
	return domain.NewUser(id), nil
}
