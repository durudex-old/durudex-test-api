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
	var images []string

	url := "https://cdn.durudex.com/attachments/" + faker.UUIDHyphenated() + "/"

	for i := 0; i < num; i++ {
		images = append(images, url+faker.Word()+".png")
	}

	return images
}
