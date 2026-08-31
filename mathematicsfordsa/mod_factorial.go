package mathematicsfordsa

const MOD = 1000000007

func BadFactoril(n int) int {
	MOD := 1000000007
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result % MOD
}

func GoodFactorial(n int) int {
	MOD := 1000000007
	result := 1

	for i := 0; i < n; i++ {
		result = (result * i) % MOD
	}
	return result
}

func AddMod(a int,b int ) int {
	return ((a % MOD) + (b % MOD)) % MOD
}

func SubMod(a int, b int) int {
	return ((a % MOD) - (b % MOD) % MOD)
}

func MulMod(a int, b int)int{
	return (( a % MOD )*(b % MOD) % MOD )
}	