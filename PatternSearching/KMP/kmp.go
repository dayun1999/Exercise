package KMP

import "unicode/utf8"

func SearchingThoughKMP(pattern string, text string) []int {
	resultIndex := make([]int, 0)
	txtLength := utf8.RuneCountInString(text)
	patLength := utf8.RuneCountInString(pattern)
	patSlice := []rune(pattern)
	txtSlice := []rune(text)
	lps := computeLPS(patSlice, patLength)
	i := 0 // index of txt[]
	j := 0 //index of pat[]
	for i < txtLength {
		if patSlice[j] == txtSlice[i] {
			i++
			j++
		}
		if j == patLength {
			resultIndex = append(resultIndex, i-j)
			j = lps[j-1]
		} else if i < txtLength && patSlice[j] != txtSlice[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}

	return resultIndex

}
func computeLPS(pattern []rune, patLength int) []int {
	lps := make([]int, patLength)
	length := 0
	i := 1
	for i < patLength {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}
