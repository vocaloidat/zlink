package base62

import (
	"fmt"
	"math"
	"strings"
)

var (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func NewBase62Tool(s string) {
	if len(s) < 62 {
		panic("长度必须达到62")
	}
	base62Chars = s
}

// EncodeIntToBase62 一个10进制转base62进制的转换器
func EncodeIntToBase62(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var result []byte
	for num > 0 {

		remainder := num % 62 // 取余数
		result = append(result, base62Chars[remainder])
		num = num / 62 // 整除取
	}

	//反转结果
	return string(reverse(result))
}

// DecodeBase62ToInt 一个base62进制转10进制的转换器
func DecodeBase62ToInt(base62Str string) uint64 {
	var result float64
	for i, char := range base62Str {
		fmt.Println("i=", i, ",char=", char)
		index := strings.IndexRune(base62Chars, char) // 获取字符在62进制中的索引
		result += float64(index) * math.Pow(62, float64(len(base62Str)-i-1))
	}
	return uint64(result)
}

// 两头反转 反转结果以获得正确的base62字符串 [1,2,3,4,5]->[5,4,3,2,1]
func reverse(result []byte) []byte {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
