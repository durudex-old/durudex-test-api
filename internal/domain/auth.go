/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Authorization tokens.
type Tokens struct {
	// JWT access token.
	Access string `json:"access"`
	// Refresh token.
	Refresh string `json:"refresh"`
}

// User Sign Up input.
type SignUpInput struct {
	// Account username.
	Username string `json:"username"`
	// User email.
	Email string `json:"email"`
	// User password.
	Password string `json:"password"`
	// User verification code.
	Code uint64 `json:"code"`
	// Client secret key.
	Secret string `json:"secret"`
}

// Validate user sign up input.
func (i SignUpInput) Validate() error {
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
	case !RxEmail.MatchString(i.Email):
		// Return invalid email graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Email",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	case i.Code == 0:
		// Return invalid code graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Code",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	default:
		return nil
	}
}

// User Sign In input.
type SignInInput struct {
	// Account login.
	Login string `json:"login"`
	// User password
	Password string `json:"password"`
	// Client secret key.
	Secret string `json:"secret"`
}

// Validate user sign in input.
func (i SignInInput) Validate() error {
	switch {
	case !RxUsername.MatchString(i.Login):
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

// Session credentials input.
type SessionCredInput struct {
	// Refresh token.
	Refresh string `json:"refresh"`
	// Client secret key.
	Secret string `json:"secret"`
}
