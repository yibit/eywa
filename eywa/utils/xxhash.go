package utils

import (
	"github.com/cespare/xxhash/v2"
)

func Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func Sum64String(str string) uint64 {
	has := xxhash.New()
	has.WriteString(str)

	return has.Sum64()
}
