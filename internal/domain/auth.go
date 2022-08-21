/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import "github.com/vektah/gqlparser/v2/gqlerror"

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
	// Account username.
	Username string `json:"username"`
	// User password
	Password string `json:"password"`
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

// Refresh token input.
type RefreshTokenInput struct {
	// Refresh token.
	Token string `json:"token"`
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
