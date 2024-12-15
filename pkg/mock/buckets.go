package mock

import (
	"math/rand"
	"strings"
)

func RandomBucketName() string {
	words := []string{
		"apple", "banana", "cherry", "date", "elderberry", "fig", "grape",
		"honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya",
		"quince", "raspberry", "strawberry", "tangerine", "ugli", "vanilla",
		"watermelon", "xigua", "yam", "zucchini",
	}

	passphrase := []string{}
	size := 6
	for i := 0; i < size; i++ {
		passphrase = append(passphrase, words[rand.Intn(len(words))])
	}

	return strings.Join(passphrase, "-")
}

func RandomObjectID() string {
	var (
		builder = strings.Builder{}
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		length  = 16
	)

	builder.Grow(length)

	for i := 0; i < length; i++ {
		builder.WriteByte(charset[rand.Intn(len(charset))])
	}

	return strings.ToLower(builder.String())
}
