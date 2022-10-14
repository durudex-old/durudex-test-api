/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/segmentio/ksuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Authorization tokens.
type Tokens struct {
	// JWT access token.
	Access string `json:"access"`
	// Refresh token.
	Refresh string `json:"refresh"`
}

// User session.
type Session struct {
	// Session id.
	Id ksuid.KSUID `json:"id"`
	// Session user id.
	UserId ksuid.KSUID `json:"userId"`
	// User session ip address.
	Ip string `json:"ip"`
	// Session expires in.
	ExpiresIn time.Time `json:"expiresIn"`
}

// Creating a new user session.
func NewSession(id ksuid.KSUID) *Session {
	return &Session{
		Id:        id,
		UserId:    ksuid.New(),
		Ip:        faker.IPv4(),
		ExpiresIn: time.Unix(faker.RandomUnixTime(), 0),
	}
}

// List of session owned by the subject.
type SessionConnection struct {
	// A list of nodes.
	Nodes []*Session `json:"nodes"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a session connection.
type SessionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Session `json:"node"`
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
	// Account username.
	Username string `json:"username"`
	// User password
	Password string `json:"password"`
	// Client secret key.
	Secret string `json:"secret"`
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

// Delete user session input.
type DeleteSessionInput struct {
	// Session id.
	Id ksuid.KSUID `json:"id"`
	// Refresh token.
	Refresh string `json:"refresh"`
	// Client secret key.
	Secret string `json:"secret"`
}

// Session credentials input.
type SessionCredInput struct {
	// Refresh token.
	Refresh string `json:"refresh"`
	// Client secret key.
	Secret string `json:"secret"`
}
