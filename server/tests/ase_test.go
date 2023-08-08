package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
	"tools/model/system"
)

func encrypt(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secret. It's common to include it at the beginning of the encrypted output.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func decrypt(cipherText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.URLEncoding.DecodeString(cipherText)
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

func TestMy(t *testing.T) {
	// 128-bit (16 bytes) key
	key := "AESKey123456789a"

	// Text to be encrypted
	plainText := "Hello, AES!"

	// Encryption
	//encryptedText, err := encrypt(plainText, key)
	//if err != nil {
	//	fmt.Println("Encryption error:", err)
	//	return
	//}
	//fmt.Println("Encrypted:", encryptedText)
	//
	//// Decryption
	//decryptedText, err := decrypt(encryptedText, key)
	//if err != nil {
	//	fmt.Println("Decryption error:", err)
	//	return
	//}
	//fmt.Println("Decrypted:", decryptedText)

	var types = "aes"
	system.InitAlgorithmMap()
	word := system.AlgorithmMap[system.Algorithm(types)](key)
	s, err := word.Encrypt(plainText)
	t.Log(s)

	s2, _ := word.Decrypt("hVa3r7NFoxwhczOG2xX20l4Y09ctfOoNmADj")
	t.Log(s2)
	if err != nil {
		t.Fatal(err)
	}
}
