package util

import (
	"math/rand"
	"time"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}
	return string(b)
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
