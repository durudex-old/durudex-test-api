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

// User session service interface.
type Session interface {
	// Creating a new fake user session.
	Create(userId ksuid.KSUID) (*domain.Tokens, error)
	// Deleting a fake user session.
	Delete(ctx context.Context, id ksuid.KSUID) (bool, error)
	// Getting a fake user session.
	Get(ctx context.Context, id ksuid.KSUID) (*domain.Session, error)
	// Getting a fake user sessions.
	GetList(ctx context.Context, sort domain.SortOptions) (*domain.SessionConnection, error)
}

// User session service structure.
type SessionService struct{ cfg *config.AuthConfig }

// Creating a new user session service.
func NewSessionService(cfg *config.AuthConfig) *SessionService {
	return &SessionService{cfg: cfg}
}

// Creating a new fake user session.
func (s *SessionService) Create(userId ksuid.KSUID) (*domain.Tokens, error) {
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

// Deleting a user session.
func (s *SessionService) Delete(ctx context.Context, id ksuid.KSUID) (bool, error) {
	if id.IsNil() {
		return false, &gqlerror.Error{
			Message:    "Session not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return true, nil
}

// Getting a user session.
func (s *SessionService) Get(ctx context.Context, id ksuid.KSUID) (*domain.Session, error) {
	if id.IsNil() {
		return nil, &gqlerror.Error{
			Message:    "Session not found",
			Extensions: map[string]interface{}{"code": domain.CodeNotFound},
		}
	}

	return domain.NewSession(id), nil
}

// Getting a user sessions.
func (s *SessionService) GetList(ctx context.Context, sort domain.SortOptions) (*domain.SessionConnection, error) {
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
