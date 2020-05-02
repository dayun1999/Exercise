package KMP

import "testing"

func TestKmpSearching(t *testing.T) {
	text := "23天科大天空大云白云天天向上abccbab"
	pattern := "ab"
	index := SearchingThoughKMP(pattern, text)
	t.Log(index)
}
