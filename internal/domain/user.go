/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// User structure.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	LastVisit time.Time `json:"lastVisit"`
	Verified  bool      `json:"verified"`
	AvatarURL *string   `json:"avatarUrl"`
}

// Creating a new user.
func NewUser(id string) *User {
	return &User{
		ID:        id,
		Username:  faker.Username(),
		LastVisit: time.Unix(faker.RandomUnixTime(), 0),
		Verified:  rand.Intn(2) == 1,
		AvatarURL: NewOptionalString("https://cdn.durudex.com/avatar/" + id + ".png"),
	}
}

func (User) IsNode() {}

// User Sign Up input.
type SignUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     uint64 `json:"code"`
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

// User forgot password input.
type ForgotPasswordInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     uint64 `json:"code"`
}

// Validate user forgot password input.
func (i ForgotPasswordInput) Validate() error {
	switch {
	case !RxEmail.MatchString(i.Email):
		// Return invalid email graphql error.
		return &gqlerror.Error{
			Message:    "Invalid Email",
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
