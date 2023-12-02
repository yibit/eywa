package encoding

import (
	"encoding/hex"
)

func HexEncode(src string) string {
	return hex.EncodeToString([]byte(src))
}

func HexDecode(src string) string {
	return hex.EncodeToString([]byte(src))
}
