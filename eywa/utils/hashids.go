package utils

import (
	hashids "github.com/speps/go-hashids/v2"
)

// https://hashids.org/go/
func HashIDs(salt string, data []int, min int) (string, error) {
	hd := hashids.NewData()

	hd.Salt = salt
	hd.MinLength = min
	if h, err := hashids.NewWithData(hd); err != nil {
		return "", err
	} else {
		return h.Encode(data)
	}
}
