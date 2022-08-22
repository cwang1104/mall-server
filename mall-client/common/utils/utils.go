package utils

import (
	"regexp"
	"strconv"
)

// 校验邮箱
//正则表达式
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//string to int32
func StrToInt32(str string) int32 {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int32(n)
}

func StrToFloat32(str string) float32 {
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return float32(n)
}
