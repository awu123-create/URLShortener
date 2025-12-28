package Utils

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 自增算法（使用数据库自增ID作为唯一数字，然后把这个数字编码成Base62字符串，即是生成的短码）
func EncodeBase62(num uint) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var result []byte
	base := uint(len(base62Chars))

	for num > 0 {
		remainder := num % base
		result = append(result, base62Chars[remainder])
		num /= base
	}

	return string(result)
}
