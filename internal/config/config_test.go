/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package config_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-test-api/internal/config"
)

// Test initialize config.
func TestConfig_Init(t *testing.T) {
	// Environment configurations.
	type env struct{ configPath, jwtSigningKey string }

	// Testing args.
	type args struct{ env env }

	// Set environments configurations.
	setEnv := func(env env) {
		os.Setenv("CONFIG_PATH", env.configPath)
		os.Setenv("JWT_SIGNING_KEY", env.jwtSigningKey)
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{
				configPath:    "fixtures/main",
				jwtSigningKey: "secret-key",
			}},
			want: &config.Config{
				Auth: config.AuthConfig{
					SigningKey: "secret-key",
					TTL:        time.Minute * 15,
				},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environments configurations.
			setEnv(tt.args.env)

			// Initialize config.
			got, err := config.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("error initialize config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error config are not similar")
			}
		})
	}
}
