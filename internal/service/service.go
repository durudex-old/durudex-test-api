/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package service

// Service structure.
type Service struct {
	Auth
	User
	Post
}

// Creating a new service.
func NewService() *Service {
	return &Service{
		Auth: NewAuthService(),
		User: NewUserService(),
		Post: NewPostService(),
	}
}
