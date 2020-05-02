package Prime

//判断一个数是否是素数
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
