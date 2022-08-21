/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package gql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/segmentio/ksuid"
)

// Marshal a KSUID to a string.
func MarshalKSUID(v ksuid.KSUID) graphql.Marshaler {
	return graphql.MarshalString(v.String())
}

// Unmarshal a KSUID from a string.
func UnmarshalKSUID(v any) (ksuid.KSUID, error) {
	switch v := v.(type) {
	case string:
		return ksuid.Parse(v)
	default:
		return ksuid.Nil, fmt.Errorf("ksuid must be a string")
	}
}
