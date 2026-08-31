package mathematicsfordsa


func BadFactoril(n int) int {
	mod := 1000000007
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result % mod
}

func GoodFactorial(n int) int {
	mod := 1000000007
	result := 1

	for i := 0; i < n; i++ {
		result = (result * i) % mod
	}
	return result
}