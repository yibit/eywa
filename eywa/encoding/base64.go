package encoding

import (
	"encoding/base64"
)

func EncodeToString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeString(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}

	return string(data)
}
