/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/segmentio/ksuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const UserCtx string = "userId"

// User type.
type User struct {
	// User id.
	Id ksuid.KSUID `json:"id"`
	// Username.
	Username string `json:"username"`
	// User last visit date.
	LastVisit time.Time `json:"lastVisit"`
	// User verified status.
	Verified bool `json:"verified"`
	// User avatar url.
	AvatarUrl *string `json:"avatarUrl"`
}

func (User) IsNode() {}

// Creating a new user.
func NewUser(id ksuid.KSUID) *User {
	return &User{
		Id:        id,
		Username:  faker.Username(),
		LastVisit: time.Unix(faker.RandomUnixTime(), 0),
		Verified:  rand.Intn(2) == 1,
		AvatarUrl: NewOptionalString("https://cdn.durudex.com/avatar/" + id.String() + ".png"),
	}
}

// Forgot user password input.
type ForgotPasswordInput struct {
	// User email.
	Email string `json:"email"`
	// New user password.
	Password string `json:"password"`
	// User verification code.
	Code uint64 `json:"code"`
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
