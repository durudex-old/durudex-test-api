/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package domain

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/segmentio/ksuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Regular fot validating strings.
const (
	Username string = "^[a-zA-Z0-9-_.]{3,40}$"
	Password string = "^[a-zA-Z0-9@$!%*?&]{8,100}$"
	Email    string = "^(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
)

var (
	RxUsername = regexp.MustCompile(Username)
	RxPassword = regexp.MustCompile(Password)
	RxEmail    = regexp.MustCompile(Email)
)

// Query sorting options.
type SortOptions struct {
	First  *int
	Last   *int
	Before *string
	After  *string
}

// Validating query sort options.
func (o SortOptions) Validate() (int, error) {
	var limit int

	// Check filter and last filters.
	switch {
	// Check if first and last filters is not nil.
	case o.First != nil && o.Last != nil:
		return 0, &gqlerror.Error{
			Message:    "Must be `first` or `last`",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	// Check if first filter is nil.
	case o.First == nil:
		// Check if last filter is nil or set last filter.
		if o.Last == nil {
			return 0, &gqlerror.Error{
				Message:    "Must be `first` or `last`",
				Extensions: map[string]interface{}{"code": CodeInvalidArgument},
			}
		} else if *o.Last > 50 || *o.Last < 1 {
			return 0, &gqlerror.Error{
				Message:    "`last` must not exceed 50 or be less than 1",
				Extensions: map[string]interface{}{"code": CodeInvalidArgument},
			}
		}

		limit = *o.Last
	// Check if first filter is nil or set last filter.
	case *o.First > 50 || *o.First < 1:
		return 0, &gqlerror.Error{
			Message:    "`first` must not exceed 50 or be less than 1",
			Extensions: map[string]interface{}{"code": CodeInvalidArgument},
		}
	// Set first filter.
	default:
		limit = *o.First
	}

	return limit, nil
}

// Creating a new random optional time.
func NewOptionalTime() *time.Time {
	if rand.Intn(2) == 1 {
		updatedAt := time.Unix(faker.RandomUnixTime(), 0)
		return &updatedAt
	} else {
		return nil
	}
}

// Creating a new random optional string.
func NewOptionalString(s string) *string {
	if rand.Intn(2) == 1 {
		return &s
	} else {
		return nil
	}
}

// Creating a new random optional array of attachments url.
func NewRandomAttachmentsURLArray(num int) []string {
	images := make([]string, num)

	url := "https://cdn.durudex.com/attachments/" + ksuid.New().String() + "/"

	for i := 0; i < num; i++ {
		images[i] = url + faker.Word() + ".png"
	}

	return images
}
