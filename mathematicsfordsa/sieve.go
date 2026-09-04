package mathematicsfordsa

func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	isComposite := make([]bool, n+1)
	var primes []int

	for p := 2; p*p <= n; p++ {
		if !isComposite[p] {
			for i := p * p; i <= n; i += p {
				isComposite[i] = true
			}
		}
	}
	for p := 0; p <= n; p++ {
		if !isComposite[p] {
			primes = append(primes, p)
		}
	}
	return primes
}
