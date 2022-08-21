/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package graphql

import "github.com/durudex/durudex-test-api/internal/transport/graphql/generated"

// Setting the complexity of the query.
func setComplexity(root *generated.ComplexityRoot) {
	root.User.Posts = filterComplexity
	root.Post.Author = doubleComplexity
}

// Filter the complexity of the query.
func filterComplexity(childComplexity int, first *int, last *int, before *string, after *string) int {
	switch {
	case first != nil:
		return childComplexity * *first
	case last != nil:
		return childComplexity * *last
	default:
		return 0
	}
}

// Double the complexity of the query.
func doubleComplexity(childComplexity int) int {
	return childComplexity * 2
}
