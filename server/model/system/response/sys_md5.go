package response

type MD5Resp struct {
	Base16Lowercase string `json:"16lowercase"` // 16小写
	Base16Uppercase string `json:"16uppercase"` // 16大写
	Base32Lowercase string `json:"32lowercase"` // 32小写
	Base32Uppercase string `json:"32uppercase"` // 32大写
}
