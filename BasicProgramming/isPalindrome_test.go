package BasicProgramming

import "testing"

func TestIsPalindrome(t *testing.T) {
	target := "helloolleh"
	result := IsPalindrome(target)
	if result{
		t.Log("是回文")
	}else{
		t.Log("不是回文")
	}
}
