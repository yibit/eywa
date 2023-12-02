package encoding

import (
	"encoding/hex"
)

func HexEncode(s string) string {
	return hex.EncodeToString([]byte(s))
}

func HexDecode(s string) string {
	data, err := hex.DecodeString(s)
	if err != nil {
		return ""
	}

	return string(data)
}
