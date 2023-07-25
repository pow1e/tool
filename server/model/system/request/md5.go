package request

// Md5Encrypt md5加密
type Md5Encrypt struct {
	Plaintext string `json:"plaintext"`
}

// Md5Decrypt md5解密
type Md5Decrypt struct {
	Ciphertext string `json:"ciphertext"`
}
