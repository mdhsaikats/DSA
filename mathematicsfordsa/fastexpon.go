package mathematicsfordsa

func Normal(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func FastExpon(a, b int) int {
	if b == 0 {
		return 1
	}
	half := FastExpon(a, b/2)
	if b%2 == 0 {
		return half * half
	}
	return a * half * half
}
