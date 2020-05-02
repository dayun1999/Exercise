package BasicProgramming

import "unicode/utf8"

//判断一个字符是否是回文
func IsPalindrome(text string) bool {
	result := true
	textSlice := []rune(text) //实现Unicode字符识别
	textLen := utf8.RuneCountInString(text)
	for i := 0; i < textLen/2; i++ {
		if textSlice[i] == textSlice[textLen-1-i] {
			continue
		}
		result = false
	}
	return result
}
