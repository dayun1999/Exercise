package BruteForce

import (
	"unicode/utf8"
)

func FindIndex(pattern, text string) int {
	//计算查找和被查找的字符串的长度
	patternLength := utf8.RuneCountInString(pattern)
	textLength := utf8.RuneCountInString(text)
	//为了方便中文字符的比对，直接转为rune
	patternRune := []rune(pattern)
	textRune := []rune(text)
	var i = 0
	for i = 0; i <= textLength-patternLength; i++ {
		j := 0
		for j = 0; j < patternLength; j++ {
			if textRune[i+j] != patternRune[j] {
				break
			}
		}
		if j == patternLength {
			return i
		}
	}
	return -1
}
