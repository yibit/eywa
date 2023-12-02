package utils

import (
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

func ULID() (id ulid.ULID, err error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	return ulid.New(ms, entropy)
}
