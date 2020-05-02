package Prime

import "testing"

func TestIsPrime(t *testing.T) {
	n := 99
	prime := IsPrime(n)
	if prime {
		t.Log("是素数")
	} else {
		t.Log("不是素数")
	}
}
