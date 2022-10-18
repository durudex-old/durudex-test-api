/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

import "github.com/durudex/durudex-test-api/internal/config"

// Service structure.
type Service struct {
	Auth    Auth
	User    User
	Session Session
	Post    Post
}

// Creating a new service.
func NewService(cfg *config.Config) *Service {
	sessionService := NewSessionService(&cfg.Auth)

	return &Service{
		Auth:    NewAuthService(sessionService, &cfg.Auth),
		User:    NewUserService(),
		Session: sessionService,
		Post:    NewPostService(),
	}
}
