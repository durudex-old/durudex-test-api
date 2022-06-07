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

	url := "https://cdn.durudex.com/attachments/" + faker.UUIDHyphenated() + "/"

	for i := 0; i < num; i++ {
		images = append(images, url+faker.Word()+".png")
	}

	return images
}
