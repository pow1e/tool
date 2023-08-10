package system

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type AesAlgorithm struct {
	Plaintext string `json:"plaintext"`
	Key       []byte `json:"key"`
}

func (a AesAlgorithm) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (a AesAlgorithm) Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(cipherBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)

	return string(cipherBytes), nil
}

func NewAESFunc(option ...string) IPassWord {
	if len(option) == 0 {
		panic("请正确填入aes加密的key")
	}
	s := option[0]
	// NewAESFunc 密钥必须是16或者24位或者32位
	bytes := fillKey([]byte(s))
	return &AesAlgorithm{
		Key: bytes,
	}
}

// 使用0填充密钥，确保密钥长度为16字节
func fillKey(key []byte) []byte {
	if len(key) >= aes.BlockSize {
		return key[:aes.BlockSize]
	}
	paddedKey := make([]byte, aes.BlockSize)
	copy(paddedKey, key)
	return paddedKey
}
