package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

const (
	AesDefaultBlockSize = 16
)

// AesSimpleEncrypt encrypts data with key using AES algorithm.
// In simple encryption mode, the user only needs to specify the key to complete the encryption.
// IV will be obtained by hashing the key. By default, PKCS7Padding and CBC modes are used.
// If key length is not 16, it will be padded with 0 or truncated to 16 bytes.
// Return empty string if error occurs.
func AesSimpleEncrypt(data, key string) string {
	key = trimByBlockSize(key)
	keyBytes := []byte(key)
	keyBytes = ZerosPadding(keyBytes, AesDefaultBlockSize)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return ""
	}

	src := PKCS7Padding([]byte(data), block.BlockSize())
	encryptData := make([]byte, len(src))
	mode := cipher.NewCBCEncrypter(block, []byte(GenIVFromKey(key)))
	mode.CryptBlocks(encryptData, src)
	return base64.StdEncoding.EncodeToString(encryptData)
}

// AesEncrypt encrypts data with key and iv using AES algorithm.
// You must make sure the length of key and iv is 16 bytes. This function does not perform any padding for key or iv.
func AesEncrypt(data, key, iv string, paddingMode PaddingMode) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	src := Padding(paddingMode, []byte(data), block.BlockSize())
	encryptData := make([]byte, len(src))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(encryptData, src)
	return base64.StdEncoding.EncodeToString(encryptData)
}

// AesSimpleDecrypt decrypts data with key using AES algorithm.
// In simple decryption mode, the user only needs to specify the key to complete the decryption.
// This function will automatically obtain the IV by hashing the key.
func AesSimpleDecrypt(data, key string) string {
	key = trimByBlockSize(key)
	keyBytes := []byte(key)
	keyBytes = ZerosPadding(keyBytes, AesDefaultBlockSize)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return ""
	}

	decodeData, _ := base64.StdEncoding.DecodeString(data)
	decryptData := make([]byte, len(decodeData))
	mode := cipher.NewCBCDecrypter(block, []byte(GenIVFromKey(key)))
	mode.CryptBlocks(decryptData, decodeData)

	original, _ := PKCS7UnPadding(decryptData)
	return string(original)
}

// AesDecrypt decrypts data with key and iv using AES algorithm.
func AesDecrypt(data, key, iv string, paddingMode PaddingMode) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	decodeData, _ := base64.StdEncoding.DecodeString(data)
	decryptData := make([]byte, len(decodeData))
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(decryptData, decodeData)

	original, _ := UnPadding(paddingMode, decryptData)
	return string(original)
}

// GenIVFromKey generates IV from key.
func GenIVFromKey(key string) (iv string) {
	hashedKey := sha256.Sum256([]byte(key))
	return trimByBlockSize(hex.EncodeToString(hashedKey[:]))
}

func trimByBlockSize(key string) string {
	if len(key) > AesDefaultBlockSize {
		return key[:AesDefaultBlockSize]
	}
	return key
}
