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
)

// User structure.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	LastVisit time.Time `json:"lastVisit"`
	Verified  bool      `json:"verified"`
	AvatarURL *string   `json:"avatarUrl"`
}

// Creating a new user.
func NewUser(id string) *User {
	return &User{
		ID:        faker.UUIDHyphenated(),
		Username:  faker.Username(),
		CreatedAt: time.Unix(faker.RandomUnixTime(), 0),
		LastVisit: time.Unix(faker.RandomUnixTime(), 0),
		Verified:  rand.Intn(2) == 1,
		AvatarURL: NewOptionalString("https://cdn.durudex.com/avatar/" + faker.UUIDHyphenated() + ".png"),
	}
}

func (User) IsNode() {}

// User forgot password input.
type ForgotPasswordInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     uint64 `json:"code"`
}
