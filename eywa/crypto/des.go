package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// EncryptDES DES encrypt
func EncryptDES(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = PKCS5Padding(data, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, key)
	ct := make([]byte, len(data))
	mode.CryptBlocks(ct, data)

	return ct, nil
}

// DecryptDES DES decrypt
func DecryptDES(crypted []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, key)
	data := make([]byte, len(crypted))
	mode.CryptBlocks(data, crypted)
	data = PKCS5UnPadding(data)

	return data, nil
}

// PKCS5Padding
func PKCS5Padding(text []byte, size int) []byte {
	padding := size - len(text)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(text, padtext...)
}

// PKCS5UnPadding
func PKCS5UnPadding(data []byte) []byte {
	length := len(data)

	padding := int(data[length-1])
	return data[:(length - padding)]
}
