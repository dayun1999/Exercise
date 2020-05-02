package BruteForce

import "testing"

func TestFindIndex(t *testing.T) {
	text := "WFASF王大云打野打枪"
	pattern := "云打"
	index := FindIndex(pattern, text)
	t.Log("索引为： ", index)
	if index != 7 {
		t.Errorf("失败，索引为%d \n", index)
	}
}
