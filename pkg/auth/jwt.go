/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package auth

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWT manager interface.
type JWT interface {
	GenerateAccessToken(subject, signingKey string, ttl time.Duration) (string, error)
	Parse(accessToken, signingKey string) (string, error)
	GenerateRefreshToken() (string, error)
}

// Generating a new jwt access token.
func GenerateAccessToken(subject, signingKey string, ttl time.Duration) (string, error) {
	// Generating a new jwt token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   subject,
	})

	return token.SignedString([]byte(signingKey))
}

// Parsing jwt access token.
func Parse(accessToken, signingKey string) (string, error) {
	// Parsing and validation token.
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (i interface{}, err error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	// Get user claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

// Generating a new refresh token.
func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
