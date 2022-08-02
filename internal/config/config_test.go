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

// Testing creating a new config.
func TestConfig_NewConfig(t *testing.T) {
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
				HTTP: config.HTTPConfig{
					Host: "api.test.durudex.com",
					Port: "8000",
					Name: "Durudex Test API",
					Cors: config.CorsConfig{
						Enable:       true,
						AllowOrigins: "*",
						AllowMethods: "*",
						AllowHeaders: "*",
					},
				},
				GraphQL: config.GraphQLConfig{ComplexityLimit: 500},
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

			// Creating a new config,
			got, err := config.NewConfig()
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
