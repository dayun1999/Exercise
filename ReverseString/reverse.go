package ReverseString

func reverseString1(name string) string {
	//将string转化为rune切片
	runes := []rune(name)
	//e.g
	//name := "郭德纲"
	//len(name) == 9 && len(runes) == 3
	for i, j := 0, len(runes)-1; i < j; i, j = j-1, i+1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
