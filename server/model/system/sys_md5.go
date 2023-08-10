package system

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

type Md5Algorithm struct {
	Plaintext string `json:"plaintext"`
}

func (m *Md5Algorithm) Encrypt(plaintext string) (string, error) {
	hash := md5.Sum([]byte(plaintext))
	encryptString := hex.EncodeToString(hash[:])
	return encryptString, nil
}

func (m *Md5Algorithm) Decrypt(ciphertext string) (string, error) {
	log.Println("实现了md5解密")
	return ciphertext, nil
}

func NewMd5Func(...string) IPassWord {
	return &Md5Algorithm{}
}
