package ReverseString_recursion

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(str string) {
	length := utf8.RuneCountInString(str)
	slice := []rune(str)
	//len("我") == 3
	if str == "" || length <= 1 {
		fmt.Print(str)
	} else {
		fmt.Print(string(slice[length-1]))
		Reverse(string(slice[0 : length-1]))
	}
}
