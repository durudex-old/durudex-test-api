/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import "github.com/99designs/gqlgen/graphql"

// GraphQL error status codes.
const (
	CodeServerError         string = "SERVER_ERROR"
	CodeInternalServerError string = "INTERNAL_SERVER_ERROR"
	CodeInvalidArgument     string = "INVALID_ARGUMENT"
	CodeNotFound            string = "NOT_FOUND"
	CodeAlreadyExists       string = "ALREADY_EXISTS"
	CodeUnauthorized        string = "UNAUTHORIZED"
)

// GraphQL Node interface.
type Node interface {
	IsNode()
}

// Upload files input.
type UploadFile struct {
	ID   int            `json:"id"`
	File graphql.Upload `json:"file"`
}
