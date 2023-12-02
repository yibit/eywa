package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKey(t *testing.T) {
	str := "aes_cbc({key=\"1234567890123456\", iv=\"1234567890123456\", padding=\"ISO78164\"})"

	assert.Equal(t, "1234567890123456", FindKey(str, "key=\"", '"'), "they should be equal")
	assert.Equal(t, "1234567890123456", FindKey(str, "iv=\"", '"'), "they should be equal")
	assert.Equal(t, "ISO78164", FindKey(str, "padding=\"", '"'), "they should be equal")

	str = "aes_cbc({\"key\"=\"1234567890123456\", \"iv\"=\"1234567890123456\", \"padding\"=\"ISO78164\"})"

	texts := FindKeys(str, []string{"key", "iv", "padding"}, "=\"", "\"", '"')
	assert.Equal(t, 3, len(texts), "they should be equal")
	assert.Equal(t, "ISO78164", texts[2], "they should be equal")
}
