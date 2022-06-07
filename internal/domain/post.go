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

// Post type.
type Post struct {
	ID          string     `json:"id"`
	Author      *User      `json:"author"`
	Text        string     `json:"text"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Attachments []string   `json:"attachments"`
}

// Creating a new post.
func NewPost(id string) *Post {
	return &Post{
		ID:          id,
		Author:      NewUser(faker.UUIDHyphenated()),
		Text:        faker.Sentence(),
		CreatedAt:   time.Unix(faker.RandomUnixTime(), 0),
		UpdatedAt:   NewOptionalTime(),
		Attachments: NewRandomAttachmentsURLArray(rand.Intn(5)),
	}
}

func (Post) IsNode() {}

// Create post input.
type CreatePostInput struct {
	AuthorID    string
	Text        string `json:"text"`
	Attachments []*UploadFile
}

// Validate create post input.
func (i CreatePostInput) Validate() error {
	if len(i.Text) > 500 {
		return &gqlerror.Error{
			Message:    "Text is too long",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	}

	return nil
}

// Update post input.
type UpdatePostInput struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Validate update post input.
func (i UpdatePostInput) Validate() error {
	if len(i.Text) > 500 {
		return &gqlerror.Error{
			Message:    "Text is too long",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	}

	return nil
}
