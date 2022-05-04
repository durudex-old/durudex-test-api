/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import "github.com/99designs/gqlgen/graphql"

// GraphQL Node interface.
type Node interface {
	IsNode()
}

// Upload files input.
type UploadFile struct {
	ID   int            `json:"id"`
	File graphql.Upload `json:"file"`
}
