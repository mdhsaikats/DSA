package mathematicsfordsa


func badFactoril(n int) int {
	mod := 1000000007
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result * mod
}