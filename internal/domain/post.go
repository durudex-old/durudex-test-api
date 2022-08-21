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

// Post type.
type Post struct {
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Post author.
	Author *User `json:"author"`
	// Post text.
	Text string `json:"text"`
	// Post updated date.
	UpdatedAt *time.Time `json:"updatedAt"`
	// Post attachments.
	Attachments []string `json:"attachments"`
}

func (Post) IsNode() {}

// Creating a new post.
func NewPost(id ksuid.KSUID) *Post {
	return &Post{
		Id:          id,
		Author:      NewUser(ksuid.New()),
		Text:        faker.Sentence(),
		UpdatedAt:   NewOptionalTime(),
		Attachments: NewRandomAttachmentsURLArray(rand.Intn(5)),
	}
}

// List of post owned by the subject.
type PostConnection struct {
	// A list of nodes.
	Nodes []*Post `json:"nodes"`
}

// An edge in a post connection.
type PostEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Post `json:"node"`
}

// Create post input.
type CreatePostInput struct {
	// Post text.
	Text string `json:"text"`
	// Post attachments.
	Attachments []*UploadFile `json:"attachments"`
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
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Post text.
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
