/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package gql_test

import (
	"bytes"
	"testing"

	"github.com/durudex/durudex-test-api/pkg/gql"
	"github.com/segmentio/ksuid"
)

// Testing marshal a KSUID to a string.
func Test_MarshalKSUID(t *testing.T) {
	// Tests structures.
	tests := []struct {
		name    string
		arg     ksuid.KSUID
		wantErr bool
	}{
		{
			name: "OK",
			arg:  ksuid.New(),
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			// Marshal a KSUID to a string.
			gql.MarshalKSUID(tt.arg).MarshalGQL(&buf)

			// Check marshal is correct.
			if len(buf.String()) != 29 {
				t.Fatalf("marshal ksuid is not correct")
			}
		})
	}
}

// Testing unmarshal a KSUID from a string.
func Test_UnmarshalKSUID(t *testing.T) {
	// Tests structures.
	tests := []struct {
		name    string
		arg     any
		wantErr bool
	}{
		{
			name: "OK",
			arg:  "2C5F4yj4xJZYlnE3DNFBQdRCCAm",
		},
		{
			name:    "Not enough characters",
			arg:     "2C5F4yj4xJZYlnE3DNFBQdR",
			wantErr: true,
		},
		{
			name:    "Too many characters",
			arg:     "2C5F4yj4xJZYlnE3DNFBQdRCCAmCCAm",
			wantErr: true,
		},
		{
			name:    "Not a string",
			arg:     2,
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal a KSUID from a string.
			got, err := gql.UnmarshalKSUID(tt.arg)
			if err != nil {
				if tt.wantErr {
					t.Skip("skipping test because of error.")
				}

				t.Fatalf("error unmarshal ksuid: %s", err.Error())
			}

			// Check ksuid is nil.
			if got.IsNil() {
				t.Fatalf("ksuid is nil")
			}
		})
	}
}
