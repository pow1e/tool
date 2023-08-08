package system

import "encoding/base64"

type Base64Algorithm struct {
	Plaintext string `json:"plaintext"`
}

func (b Base64Algorithm) Encrypt(plaintext string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(plaintext)), nil
}

func (b Base64Algorithm) Decrypt(ciphertext string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(ciphertext)
	return string(decodeString), err
}

func NewBase64Func(...string) IPassWord {
	return &Base64Algorithm{}
}
