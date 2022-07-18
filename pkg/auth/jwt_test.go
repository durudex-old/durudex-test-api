/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package auth_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-test-api/pkg/auth"
)

// Testing generating a new jwt access token.
func Test_GenerateAccessToken(t *testing.T) {
	// Testing args.
	type args struct {
		subject    string
		signingKey string
		ttl        time.Duration
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				subject:    "1",
				signingKey: "secret-key",
				ttl:        time.Hour * 9999,
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a new jwt access token.
			got, err := auth.GenerateAccessToken(tt.args.subject, tt.args.signingKey, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("error generating access token: %s", err.Error())
			}

			// Check access token is empty.
			if got == "" {
				t.Error("error access token is empty")
			}
		})
	}
}

// Testing generating a new refresh token.
func Test_GenerateRefreshToken(t *testing.T) {
	// Tests structures.
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "OK"},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a new refresh token.
			got, err := auth.GenerateRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("error generating refresh token: %s", err.Error())
			}

			// Check refresh token is empty.
			if got == "" {
				t.Error("error refresh token is empty")
			}
		})
	}
}

// Testing parsing jwt access token.
func Test_Parse(t *testing.T) {
	// Testing args.
	type args struct{ accessToken, signingKey string }

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQzMzI3NjMsInN1YiI6IjEifQ.xEKQVpR4-IGc13wz43LN0TeDfXhBbX57Qe_DVloyJvM",
				signingKey:  "super-key",
			},
			want: "1",
		},
		{
			name:    "Invalid Access Token",
			args:    args{accessToken: "", signingKey: "super-key"},
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parsing jwt access token
			got, err := auth.Parse(tt.args.accessToken, tt.args.signingKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("error parsing access token: %s", err.Error())
			}

			// Check for similarity of a claims.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error claims are not similar: %s", err.Error())
			}
		})
	}
}
