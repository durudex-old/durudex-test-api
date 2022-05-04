/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

// Authorization tokens.
type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// User Sign Up input.
type SignUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     uint64 `json:"code"`
}

// User Sign In input.
type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string
}

// Refresh tokens input.
type RefreshTokenInput struct {
	Token string `json:"token"`
	IP    string
}
