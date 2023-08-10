package tests

// types_test.go
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

// IPassWord 密码接口
type IPassWord interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	IsSetKey() bool
}

type NewPasswordFunc func(option ...string) IPassWord

var TypeMap map[Algorithm]NewPasswordFunc

type Algorithm string

const (
	MD5    Algorithm = "md5"
	AES    Algorithm = "ase"
	RSA    Algorithm = "rsa"
	Base64 Algorithm = "base64"
)

type Md5Algorithm struct {
	Plaintext string `json:"plaintext"`
}

func (m *Md5Algorithm) Encrypt(plaintext string) (string, error) {
	log.Println("实现了md5加密")
	hash := md5.Sum([]byte(plaintext))
	encryptString := hex.EncodeToString(hash[:])
	return encryptString, nil
}

func (m *Md5Algorithm) Decrypt(ciphertext string) (string, error) {
	log.Println("实现了md5解密")
	return "", nil
}

func (m *Md5Algorithm) IsSetKey() bool {
	return false
}

type AesAlgorithm struct {
	Plaintext string `json:"plaintext"`
	Key       []byte `json:"key"`
}

func (a AesAlgorithm) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(a.Key)
	log.Println("a的key是", a.Key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len([]byte(plaintext)))
	iv := ciphertext[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (a AesAlgorithm) Decrypt(ciphertext string) (string, error) {
	c := []byte(ciphertext)
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	if len(c) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := c[:aes.BlockSize]
	c = c[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(c, c)

	return string(c), nil
}

func (a AesAlgorithm) IsSetKey() bool {
	return true
}

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

func (b Base64Algorithm) IsSetKey() bool {
	return false
}

func NewMd5Func(...string) IPassWord {
	return &Md5Algorithm{}
}

func NewBase64Func(...string) IPassWord {
	return &Base64Algorithm{}
}

func NewAESFunc(k ...string) IPassWord {
	// NewAESFunc 密钥必须是16或者24位或者32位
	key := k[0]
	return &AesAlgorithm{
		Key: fillKey([]byte(key)),
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

func TestType(t *testing.T) {
	// 初始化map
	TypeMap = make(map[Algorithm]NewPasswordFunc)
	TypeMap[MD5] = NewMd5Func
	TypeMap[Base64] = NewBase64Func
	TypeMap[AES] = NewAESFunc

	var types string = "ase"
	var plaintext string = "Hello, AES!"
	algorithm := Algorithm(types)
	passwordFunc := TypeMap[algorithm]("AESKey123456789a")
	//encrypt, err := passwordFunc.Encrypt(plaintext)
	//decrypt, err := passwordFunc.Decrypt(encrypt)
	if ok := passwordFunc.IsSetKey(); ok {
	}
	encrypt, err := passwordFunc.Encrypt(plaintext)
	// TODO 这里解密应该传入密文原来的没有base64编码的值
	decrypt, err := passwordFunc.Decrypt(encrypt)
	t.Log(encrypt)
	t.Log(decrypt)
	t.Log(err)

}
