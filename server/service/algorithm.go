package service

import (
	"fmt"
	"tools/model/system"
	resp "tools/model/system/response"
	"tools/pkg/utils"
)

func (a *AlgorithmSrv) Encrypt(req *system.AlgorithmReq) (map[system.Algorithm]interface{}, error) {
	m := make(map[system.Algorithm]interface{})
	for algorithm, newFunc := range system.AlgorithmMap {
		passFunc := newFunc(req.Key)
		switch algorithm {
		case system.MD5:
			data := utils.GetMD5(req.Plaintext)
			m[algorithm] = resp.MD5Resp{Base16Lowercase: data[0],
				Base16Uppercase: data[1],
				Base32Lowercase: data[2],
				Base32Uppercase: data[3],
			}
		default:
			encrypt, err := passFunc.Encrypt(req.Plaintext)
			if err != nil {
				return nil, fmt.Errorf("%s算法加密失败", algorithm)
			}
			m[algorithm] = encrypt
		}
	}
	return m, nil
}
