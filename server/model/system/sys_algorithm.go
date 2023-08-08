package system

// IPassWord 密码接口
type IPassWord interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type AlgorithmReq struct {
	Plaintext string `json:"plaintext"`
	Key       string `json:"key"`
}

type NewPasswordFunc func(opt ...string) IPassWord

type Algorithm string

const (
	MD5    Algorithm = "md5"
	AES    Algorithm = "aes"
	RSA    Algorithm = "rsa"
	Base64 Algorithm = "base64"
)

var AlgorithmMap map[Algorithm]NewPasswordFunc

func InitAlgorithmMap() {
	// 初始化map
	AlgorithmMap = make(map[Algorithm]NewPasswordFunc)
	AlgorithmMap[MD5] = NewMd5Func
	AlgorithmMap[Base64] = NewBase64Func
	AlgorithmMap[AES] = NewAESFunc
}
