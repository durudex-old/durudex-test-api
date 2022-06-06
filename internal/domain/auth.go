/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import "github.com/vektah/gqlparser/v2/gqlerror"

// Authorization tokens.
type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// User Sign In input.
type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string
}

// Validate user sign in input.
func (i SignInInput) Validate() error {
	switch {
	case !RxUsername.MatchString(i.Username):
		// Return invalid username graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Username",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	case !RxPassword.MatchString(i.Password):
		// Return invalid password graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Password",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	default:
		return nil
	}
}

// Refresh tokens input.
type RefreshTokenInput struct {
	Token string `json:"token"`
	IP    string
}

// Validate refresh tokens input.
func (i RefreshTokenInput) Validate() error {
	switch {
	case i.Token == "":
		// Return invalid token graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Access Token",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	default:
		return nil
	}
}
