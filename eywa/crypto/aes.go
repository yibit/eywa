package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"eywa/utils"
)

// EncryptAES
func EncryptAES(text string, key string, iv []byte, padding string) []byte {
	data, k := []byte(text), []byte(key)

	// k 16 - 24 - 32
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil
	}

	size := block.BlockSize()

	if padding == "ISO78164" {
		data = ISO78164Padding(data, size)
	} else {
		data = PKCS7Padding(data, size)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	ct := make([]byte, len(data))
	mode.CryptBlocks(ct, data)

	return ct
}

// DecryptAES
func DecryptAES(cryted []byte, key string, iv []byte, padding string) []byte {
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	data := make([]byte, len(cryted))
	mode.CryptBlocks(data, cryted)

	if padding == "ISO78164" {
		data = ISO78164UnPadding(data)
	} else {
		data = PKCS7UnPadding(data)
	}

	return data
}

// ISO78164Padding ISO 781-64 Padding
func ISO78164Padding(text []byte, size int) []byte {
	pad_len := size - len(text)%size
	pad_text := bytes.Repeat([]byte{byte(0x00)}, pad_len)
	pad_text[0] = 0x80

	return append(text, pad_text...)
}

// ISO78164UnPadding ISO 781-64 UnPadding
func ISO78164UnPadding(data []byte) []byte {
	if len(data) <= 0 {
		return data
	}

	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == 0x00 {
			continue
		}
		if data[i] == 0x80 {
			return data[:i]
		}
	}

	return data
}

// GetAES_IV get a IV for AES
func GetAES_IV() []byte {
	return utils.UUID2Byte()
}

// PKCS7Padding padding PKCS7
func PKCS7Padding(text []byte, size int) []byte {
	padding := size - len(text)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(text, padtext...)
}

// PKCS7UnPadding unpadding PKCS7
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	padding := int(data[length-1])

	return data[:(length - padding)]
}
